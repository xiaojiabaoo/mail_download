package download_ser

import (
	"fmt"
	"github.com/emersion/go-imap"
	goImapId "github.com/emersion/go-imap-id"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/charset"
	"github.com/emersion/go-message/mail"
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
		err      error
		done     = make(chan error, 1)
		seqset   = new(imap.SeqSet)
		clients  = &client.Client{}
		messages = make(chan *imap.Message, 10)
		section  = imap.BodySectionName{}
		items    = []imap.FetchItem{section.FetchItem()}
		mbox     = &imap.MailboxStatus{}
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
		fmt.Println("获取邮件箱出现错误：" + err.Error())
		return customErr.New(customErr.SYSTEM_ERROR, "")
	}
	seqset.AddRange(1, mbox.Messages)
	done = make(chan error, 1)
	go func() {
		done <- clients.Fetch(seqset, items, messages)
	}()
	imap.CharsetReader = charset.Reader
	return body(messages, param.Url, param.Time)
}

func body(messages chan *imap.Message, url, times string) error {
	var index = 1
	for msg := range messages {
		var (
			err     error
			mr      *mail.Reader
			subject string
			date    time.Time
			section = imap.BodySectionName{}
		)
		text := msg.GetBody(&section)
		if text != nil {
			mr, err = mail.CreateReader(text)
			if err != nil {
				fmt.Println("创建邮件内容信息对象错误：" + err.Error())
				return customErr.New(customErr.SYSTEM_ERROR, "")
			}
			date, _ = mr.Header.Date()
			subject, _ = mr.Header.Subject()
			if !strings.Contains(subject, "INVOICE") || strings.Contains(subject, "回复") {
				continue
			}
			if !strings.Contains(subject, "#") && !strings.Contains(subject, "/") {
				continue
			}
			if times != "" {
				dateTime := tools.StrToDateTime(times)
				unix := dateTime.Unix()
				if date.Unix() < unix {
					continue
				}
			}
			fmt.Println(fmt.Sprintf(`%d--发送日期：%s，--主题：%s`, index, date.Format("2006-01-02 15:04:05"), subject))
			index++
			// 处理邮件正文
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
					fmt.Println("获取邮件部分正文出现错误，错误信息：" + err.Error())
				}
				if part != nil {
					switch header := part.Header.(type) {
					case *mail.AttachmentHeader: // 附件
						filename, err = header.Filename()
						if err == nil && filename != "" {
							content, err = ioutil.ReadAll(part.Body)
							if err != nil {
								fmt.Println("读取邮件中的附件出现错误，邮件发送日期：" + date.Format("2006-01-02 15:04:05") + "；邮件主题：" + subject + "；附件名称：" + filename + "；错误信息：" + err.Error())
							} else {
								if strings.Contains(subject, "#") {
									split := strings.Split(subject, "#")
									filename = split[1]
								} else if strings.Contains(subject, "/") {
									split := strings.Split(subject, "/")
									filename = split[1]
								}
								filename = strings.TrimSpace(filename)
								if filename == "" {
									filename = tools.CreateUuid()
								}
								err = writeFile(filename+".pdf", url, content, date)
								if err != nil {
									return err
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("------------------------------------所有邮件下载完成------------------------------------")
	return nil
}

func writeFile(filename, url string, content []byte, date time.Time) error {
	var (
		file = &os.File{}
		err  error
	)
	url = url + "\\" + date.Format("2006-01-02")
	//检测目录是否存在
	if !tools.Exists(url) {
		//创建目录
		err = os.MkdirAll(url, os.ModePerm)
		if err != nil {
			fmt.Println("创建目录出现错误：" + err.Error())
			return customErr.New(customErr.SYSTEM_ERROR, "")
		}
	}
	//如果这个文件存在了，不在保存该同名文件
	if tools.Exists(url + "\\" + filename) {
		return nil
	}
	file, err = os.Create(url + "\\" + filename)
	if err != nil {
		fmt.Println("创建文件出现错误：" + err.Error())
		return customErr.New(customErr.SYSTEM_ERROR, "")
	}
	_, err = io.WriteString(file, string(content)) //写入文件(字符串)
	if err != nil {
		fmt.Println("写入文件出错：" + err.Error())
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
