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
	tools.ProcessMap.Store(param.Account, param.Serial) // 添加本次操作表示
	tools.Logger(param.Serial, "-----------------------分隔符-开始准备工作------------------------", "")
	tools.Logger(param.Serial, fmt.Sprintf(`客户端发出下载邮箱附件的指令，开始执行指令；请求信息：%+v`, param), "")
	tools.Logger(param.Serial, "开始模拟登录邮箱", "")
	clients, err = loginEmail(param.Server, param.Account, param.Password)
	if err != nil {
		tools.ProcessMap.Delete(param.Account)
		return customErr.New(customErr.EMAIL_LOGIN_ERROR, "")
	}
	tools.Logger(param.Serial, "模拟登录邮箱成功", "")
	go body(clients, param)
	return nil
}

func body(clients *client.Client, param request_model.DownloadParam) error {
	var (
		total, number, count int // total.邮件总数.包含下载的和不符合规则后被跳过的数量；count.下载的邮件数量
		done                 = make(chan error, 1)
		seqset               = new(imap.SeqSet)
		section              = imap.BodySectionName{}
		items                = []imap.FetchItem{section.FetchItem()}
		mbox                 = &imap.MailboxStatus{}
		err                  error
	)
	defer func() {
		if r := recover(); r != nil {
			tools.ProcessMap.Delete(param.Account)
			sendMessage(param, 0, 0, 0)
		}
	}()
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
		tools.ProcessMap.Delete(param.Account)
		return customErr.New(customErr.SYSTEM_ERROR, "")
	}
	tools.Logger(param.Serial, fmt.Sprintf(`开始获取邮箱中的所有邮件`), "")
	var lens = mbox.Messages
	var start uint32 = 1
	if param.Count > 0 {
		if param.Count < mbox.Messages {
			start = mbox.Messages - (param.Count - 1)
			lens = param.Count - 1
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
	tools.Logger(param.Serial, fmt.Sprintf(`已获取到邮箱中的邮件，开始按照条件筛选合适的附件`), "")
	for msg := range messages {
		var (
			mr       *mail.Reader
			subject  string
			date     time.Time
			sections = imap.BodySectionName{}
		)
		total++
		tools.Logger(param.Serial, fmt.Sprintf(`-----------------------分隔符-开始操作第%d个邮件（邮箱中所有的邮件，不仅限于符合操作条件的）------------------------`, total), "")
		text := msg.GetBody(&sections)
		if text != nil {
			mr, err = mail.CreateReader(text)
			if err != nil {
				tools.Logger(param.Serial, fmt.Sprintf(`创建邮件Reader对象出现错误：%s`, err.Error()), tools.LOG_LEVEL_SYSTEM_ERROR)
				continue
			}
			date, _ = mr.Header.Date()
			subject, _ = mr.Header.Subject()
			if !strings.Contains(subject, "INVOICE") || strings.Contains(subject, "回复") ||
				strings.Contains(subject, "RE:") || strings.Contains(subject, "Re: ") {
				tools.Logger(param.Serial, fmt.Sprintf(`检测到发送时间：%s，主题：%s，的邮件不符合指定的规则，已跳过`, date.Format("2006-01-02 15:04:05"), subject), tools.LOG_LEVEL_ERROR)
				number++
				continue
			}
			if !strings.Contains(subject, "#") && !strings.Contains(subject, "/") {
				tools.Logger(param.Serial, fmt.Sprintf(`检测到发送时间：%s，主题：%s，的邮件不符合指定的规则，已跳过`, date.Format("2006-01-02 15:04:05"), subject), tools.LOG_LEVEL_ERROR)
				number++
				continue
			}
			if param.Time != "" {
				dateTime := tools.StrToDateTime(param.Time)
				unix := dateTime.Unix()
				if date.Unix() < unix {
					number++
					continue
				}
			}
			tools.Logger(param.Serial, fmt.Sprintf(`当前操作邮件的发送日期：%s，主题：%s`, date.Format("2006-01-02 15:04:05"), subject), "")
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
					tools.Logger(param.Serial, fmt.Sprintf(`获取当前邮件部分正文出现错误：%s`, err.Error()), tools.LOG_LEVEL_ERROR)
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
								tools.Logger(param.Serial, fmt.Sprintf(`读取当前邮件中的附件出现错误，附件名称：%s，错误信息：%s`, filename, err.Error()), tools.LOG_LEVEL_ERROR)
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
									tools.Logger(param.Serial, fmt.Sprintf(`把附件中的信息下载写入PDF息错误，附件名称：%s，错误信息：%s`, filename, err.Error()), tools.LOG_LEVEL_ERROR)
									continue
								}
								tools.Logger(param.Serial, fmt.Sprintf(`邮件附件下载成功：附件名称：%s，邮件主题：%s，邮件发送时间：%s`, filename, subject, date.Format("2006-01-02 15:04:05")), "")
								number++
								count++
							}
						}
					}
				}
			}
		}
	}
	tools.Logger(param.Serial, fmt.Sprintf(`邮件附件下载结束：本次共获取到了%d封邮件；其中下载成功%d封`, total, count), "")
	//邮件通知
	sendMessage(param, total, number, count)
	tools.ProcessMap.Delete(param.Account)
	return nil
}

func sendMessage(param request_model.DownloadParam, total, number, count int) {
	if param.Inform == "on" {
		var (
			account = param.Account
			result  string
		)
		if param.InformAccount != "" {
			account = param.InformAccount
		}
		switch {
		case number == 0:
			result = tools.CCL_RESULT_FAIL
		case total == number && count > 0:
			result = tools.CCL_RESULT_ALL_SUCCESS
		case count == 0:
			result = tools.CCL_RESULT_ALL_FAIL
		default:
			result = tools.CCL_RESULT_PART_FAIL
		}
		tools.MailAttachment(account, result, param.Serial)
	}
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
