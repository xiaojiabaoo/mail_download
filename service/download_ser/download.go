package download_ser

import (
	"fmt"
	"github.com/emersion/go-imap"
	goImapId "github.com/emersion/go-imap-id"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/charset"
	"github.com/emersion/go-message/mail"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"mail_download/model/request_model"
	"mail_download/tools"
	customErr "mail_download/tools/error"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
)

func Download(param request_model.DownloadParam) error {
	var (
		err     error
		clients = &client.Client{}
		now     = time.Now()
	)
	//校验参数
	err = checkParam(&param)
	if err != nil {
		return err
	}
	//检测目录是否存在
	if tools.Exists(param.Url) {
		//检测是否为目录，如果是文件，则打回操作
		if !tools.IsDir(param.Url) {
			return customErr.New(customErr.DOWNLOAD_URL_ERROR, "")
		}
	}
	clients, err = loginEmail(param.Server, param.Account, param.Password)
	if err != nil {
		fmt.Println("登录邮箱出错：" + err.Error())
		return customErr.New(customErr.EMAIL_LOGIN_ERROR, "")
	}
	go body(clients, param, now)
	return nil
}

func body(clients *client.Client, param request_model.DownloadParam, now time.Time) error {
	var (
		index    = 1
		errorMsg []string
		texts    string
		minute   float64
		done     = make(chan error, 1)
		seqset   = new(imap.SeqSet)
		section  = imap.BodySectionName{}
		items    = []imap.FetchItem{section.FetchItem()}
		mbox     = &imap.MailboxStatus{}
		err      error
	)
	newClient := goImapId.NewClient(clients)
	newClient.ID(
		goImapId.ID{
			goImapId.FieldName:    "IMAPClient",
			goImapId.FieldVersion: "2.1.0",
		},
	)
	defer clients.Close()
	mbox, err = clients.Select("INBOX", false)
	if err != nil {
		fmt.Println("程序已终止，获取邮件箱出现错误：" + err.Error())
		return customErr.New(customErr.SYSTEM_ERROR, "")
	}
	fmt.Println("正在解析加载中，过程可能会消耗较长的时间，建议您最小化当前页面，继续您的其他工作！（不影响本程序执行）")
	var lens = mbox.Messages
	var start uint32 = 1
	if param.Count > 0 {
		if param.Count < mbox.Messages {
			start = mbox.Messages - param.Count
			lens = param.Count
		}
	} else {
		start, lens = checkMailCount(clients, lens, lens, param.Time)
	}
	seqset.AddRange(start, mbox.Messages)
	done = make(chan error, lens)
	messages := make(chan *imap.Message, lens)
	go func() {
		done <- clients.Fetch(seqset, items, messages)
	}()
	imap.CharsetReader = charset.Reader
	fmt.Println("已获取到邮件，正在按条件筛选合适的附件中......")
	for msg := range messages {
		var (
			mr       *mail.Reader
			subject  string
			date     time.Time
			sections = imap.BodySectionName{}
		)
		text := msg.GetBody(&sections)
		if text != nil {
			mr, err = mail.CreateReader(text)
			if err != nil {
				errorMsg = append(errorMsg, "创建邮件内容信息对象出现错误："+err.Error()+"；上一封邮件信息："+texts)
				continue
			}
			date, _ = mr.Header.Date()
			subject, _ = mr.Header.Subject()
			if !strings.Contains(subject, "INVOICE") || strings.Contains(subject, "回复") ||
				strings.Contains(subject, "RE:") || strings.Contains(subject, "Re: ") {
				continue
			}
			if !strings.Contains(subject, "#") && !strings.Contains(subject, "/") {
				continue
			}
			if param.Time != "" {
				dateTime := tools.StrToDateTime(param.Time)
				unix := dateTime.Unix()
				if date.Unix() < unix {
					continue
				}
			}
			texts = fmt.Sprintf(`正在下载第%d封邮件附件，发送日期：%s，主题：%s`, index, date.Format("2006-01-02 15:04:05"), subject)
			fmt.Println(texts)
			index++
			// 处理邮件正文
			var fileNum uint
			for {
				var (
					part     *mail.Part
					filename string
					content  []byte
				)
				part, err = mr.NextPart()
				if err == io.EOF {
					break
				} else if err != nil {
					errorMsg = append(errorMsg, "获取邮件部分正文出现错误，错误信息："+err.Error()+"；上一封邮件信息："+texts)
				}
				if part != nil {
					switch header := part.Header.(type) {
					case *mail.AttachmentHeader: // 附件
						filename, err = header.Filename()
						if err == nil && filename != "" {
							fileNum++
							if fileNum >= 2 {
								continue
							}
							content, err = ioutil.ReadAll(part.Body)
							if err != nil {
								errorMsg = append(errorMsg, "读取邮件中的附件出现错误，邮件发送日期："+date.Format("2006-01-02 15:04:05")+"；邮件主题："+subject+"；附件名称："+filename+"；错误信息："+err.Error())
							} else {
								if strings.Contains(subject, "#") {
									i := strings.Index(subject, "#")
									filename = strings.ReplaceAll(subject[i+1:], "/", "和")
								} else if strings.Contains(subject, "/") {
									i := strings.Index(subject, "/")
									filename = strings.ReplaceAll(subject[i+1:], "/", "和")
								}
								filename = strings.TrimSpace(filename)
								if filename == "" {
									filename = tools.CreateUuid()
								}
								err = writeFile(filename, param, content, date)
								if err != nil {
									errorMsg = append(errorMsg, "发送日期："+date.Format("2006-01-02 15:04:05")+"；邮件主题："+subject+"；附件名称："+filename+"；错误信息："+err.Error())
									continue
								}
							}
						}
					}
				}
			}
		}
	}
	minute = (float64(time.Now().Unix()) - float64(now.Unix())) / 60
	fmt.Println(fmt.Sprintf(`------------------------------------所有邮件下载完成，总耗时：%v分钟------------------------------------`, minute))
	if len(errorMsg) >= 1 {
		fmt.Println("------------------------------------以下是下载附件时记录的报错信息-开始------------------------------------")
		for k, v := range errorMsg {
			fmt.Println(fmt.Sprintf("--%d:", k+1) + v)
		}
		fmt.Println("------------------------------------以上是下载附件时记录的报错信息-结束------------------------------------")
	}
	//邮件通知
	if param.Inform == "on" {
		fmt.Println("准备发送通知邮件，准备数据中......")
		var (
			subject, mailBody string
			account           = param.Account
		)
		if param.InformAccount != "" {
			account = param.InformAccount
		}
		if index > 1 {
			index = index - 1
		}
		switch {
		case len(errorMsg) == index || (index == 1 && len(errorMsg) > 0):
			subject = "全部下载出错"
		case len(errorMsg) > 0 && index > 1:
			subject = "部分下载出错"
		case len(errorMsg) == 0:
			subject = "全部下载成功"
		}
		mailBody = fmt.Sprintf(`本次共下载了%d个附件，错误%d个；`, index, len(errorMsg))
		for k, v := range errorMsg {
			if k == 0 {
				mailBody += fmt.Sprintf("未下载成功的错误邮件：\n")
			}
			mailBody += v + "\n"
		}

		err = tools.SendMail([]string{account}, mailBody, subject)
		if err != nil {
			fmt.Println(fmt.Sprintf("发送通知邮件失败，发送账号：%s，失败原因：%s", account, err.Error()))
		} else {
			fmt.Println(fmt.Sprintf("发送通知邮件完成，发送账号：%s，请前往邮箱中查看", account))
		}
	}
	return nil
}

func checkMailCount(clients *client.Client, start, count uint32, time string) (uint32, uint32) {
	var (
		stop    uint32
		seqset  = new(imap.SeqSet)
		done    = make(chan error, 1)
		section = imap.BodySectionName{}
		items   = []imap.FetchItem{section.FetchItem()}
		goOn    bool
		index   int
	)
	if start > 200 {
		start = start - 200
		stop = count - start
	}
	seqset.AddRange(start, count)
	done = make(chan error, stop)
	messages := make(chan *imap.Message, stop)
	go func() {
		done <- clients.Fetch(seqset, items, messages)
	}()
	imap.CharsetReader = charset.Reader
	for msg := range messages {
		var (
			sections = imap.BodySectionName{}
		)
		if index >= 1 {
			break
		}
		index++
		text := msg.GetBody(&sections)
		if text == nil {
			continue
		}
		mr, err := mail.CreateReader(text)
		if err != nil {
			continue
		}
		date, _ := mr.Header.Date()
		unix := tools.StrToDateTime(time).Unix()
		if date.Unix() >= unix {
			goOn = true
		}
	}
	//递归
	if goOn {
		start, stop = checkMailCount(clients, start, count, time)
	}
	return start, stop
}

func writeFile(filename string, param request_model.DownloadParam, content []byte, date time.Time) error {
	var (
		file = &os.File{}
		err  error
		now  = time.Now()
	)
	switch param.Classify {
	case "year":
		param.Url = param.Url + "\\" + fmt.Sprintf(`%d年`, date.Year())
	case "month":
		param.Url = param.Url + "\\" + fmt.Sprintf(`%d年`, date.Year()) + "\\" + fmt.Sprintf(`%d月`, date.Month())
	case "day":
		param.Url = param.Url + "\\" + fmt.Sprintf(`%d年`, date.Year()) + "\\" + fmt.Sprintf(`%d月`, date.Month()) + "\\" + fmt.Sprintf(`%d日`, date.Day())
	default:
		return customErr.New(customErr.FILE_DOWNLOAD_CLASSIFY_ERROR, "")
	}
	//检测目录是否存在
	if !tools.Exists(param.Url) {
		//创建目录
		err = os.MkdirAll(param.Url, os.ModePerm)
		if err != nil {
			return errors.New("创建目录出现错误：" + err.Error())
		}
	}
	//如果这个文件存在了，不在保存该同名文件
	files := param.Url + "\\" + filename + ".pdf"
	if tools.Exists(files) {
		if param.Type == "jump" {
			return nil
		} else if param.Type == "all_reserved" {
			files = fmt.Sprintf(`%s_%d.pdf`, param.Url+"\\"+filename+"-副本", now.Unix())
		}
	}
	file, err = os.Create(files)
	if err != nil {
		return errors.New("创建文件出现错误：" + err.Error())
	}
	_, err = io.WriteString(file, string(content)) //写入文件(字符串)
	if err != nil {
		return errors.New("写入文件出错：" + err.Error())
	}
	return nil
}

func loginEmail(server, user, pass string) (*client.Client, error) {
	var (
		clients = &client.Client{}
		dial    = new(net.Dialer)
		err     error
	)
	dial.Timeout = time.Duration(3) * time.Second
	clients, err = client.DialWithDialerTLS(dial, server, nil)
	if err != nil {
		return clients, err
	}
	err = clients.Login(user, pass)
	if err != nil {
		return clients, err
	}
	return clients, nil
}

func checkParam(param *request_model.DownloadParam) error {
	var (
		leapMonthMap         = map[int]struct{}{1: {}, 3: {}, 5: {}, 7: {}, 8: {}, 10: {}, 12: {}}
		plainMonthMap        = map[int]struct{}{4: {}, 6: {}, 9: {}, 11: {}}
		dateTime             []string
		date, times          string
		month, day           string
		hour, minute, second string
		leap, plain          bool
	)

	if param.Server == "" {
		return customErr.New(customErr.SERVER_ADDR_ERROR, "")
	}
	if param.Account == "" {
		return customErr.New(customErr.EMAIL_ACCOUNT_ERROR, "")
	}
	if param.Password == "" {
		return customErr.New(customErr.EMAIL_PASSWORD_ERROR, "")
	}
	if param.Type != "cover" && param.Type != "jump" && param.Type != "all_reserved" {
		return customErr.New(customErr.FILE_TYPE_EMPTY_ERROR, "")
	}
	if param.Classify != "year" && param.Classify != "month" && param.Classify != "day" {
		return customErr.New(customErr.FILE_DOWNLOAD_CLASSIFY_ERROR, "")
	}
	if param.Url == "" {
		param.Url = "C:\\mail_file_download"
	}
	if param.Count > 0 && param.Time != "" {
		return customErr.New(customErr.TIME_COUNT_ERROR, "")
	}
	if param.Count == 0 && param.Time == "" {
		return customErr.New(customErr.TIME_COUNT_ERROR, "")
	}
	if param.Time != "" {
		match, _ := regexp.MatchString("\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}", param.Time)
		if !match {
			return customErr.New(customErr.DOWNLOAD_TIME_ERROR, "")
		}
		dateTime = strings.Split(param.Time, " ")
		if len(dateTime) != 2 {
			return customErr.New(customErr.DOWNLOAD_TIME_ERROR, "")
		}
		date = dateTime[0]
		times = dateTime[1]
		dateTime = strings.Split(date, "-")
		if len(dateTime) != 3 {
			return customErr.New(customErr.DOWNLOAD_TIME_ERROR, "")
		}
		month = dateTime[1]
		day = dateTime[2]
		dateTime = strings.Split(times, ":")
		if len(dateTime) != 3 {
			return customErr.New(customErr.DOWNLOAD_TIME_ERROR, "")
		}
		hour = dateTime[0]
		minute = dateTime[1]
		second = dateTime[2]
		_, leap = leapMonthMap[tools.Str2Int(month)]
		_, plain = plainMonthMap[tools.Str2Int(month)]
		switch {
		case tools.Str2Int(month) > 12:
			return customErr.New(customErr.DOWNLOAD_TIME_ERROR, "")
		case tools.Str2Int(month) == 2:
			year := time.Now().Year()
			if isLeapYear(year) {
				if tools.Str2Int(day) > 29 {
					return customErr.New(customErr.DOWNLOAD_TIME_ERROR, "")
				}
			} else {
				if tools.Str2Int(day) > 28 {
					return customErr.New(customErr.DOWNLOAD_TIME_ERROR, "")
				}
			}
		case leap:
			if tools.Str2Int(day) > 31 {
				return customErr.New(customErr.DOWNLOAD_TIME_ERROR, "")
			}
		case plain:
			if tools.Str2Int(day) > 30 {
				return customErr.New(customErr.DOWNLOAD_TIME_ERROR, "")
			}
		case tools.Str2Int(hour) > 23:
			return customErr.New(customErr.DOWNLOAD_TIME_ERROR, "")
		case tools.Str2Int(minute) > 60:
			return customErr.New(customErr.DOWNLOAD_TIME_ERROR, "")
		case tools.Str2Int(second) > 60 || len(second) == 3:
			return customErr.New(customErr.DOWNLOAD_TIME_ERROR, "")
		}
	}
	return nil
}

// 判断是否为闰年
func isLeapYear(year int) bool {
	// 闰年的条件是：能被4整除但不能被100整除，或者能被400整除
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
