package download_ser

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"mail_download/model"
	"mail_download/model/request_model"
	"mail_download/model/response_model"
	"mail_download/service/excel_ser"
	"mail_download/tools"
	customErr "mail_download/tools/error"
	"os"
	"sort"
	"strings"
	"time"
)

func Download(param request_model.DownloadParam) error {
	var (
		err     error
		clients = &client.Client{}
	)
	//校验参数
	if err = CheckParam(&param); err != nil {
		return err
	}
	//检测目录是否存在、是否为目录；如果是文件，则打回操作
	if tools.Exists(param.Url) && !tools.IsDir(param.Url) {
		return customErr.New(customErr.DOWNLOAD_URL_ERROR, "")
	}
	tools.ProcessMap.Store(param.Account, param.Serial) // 添加本次操作表示
	tools.Logger(param.Serial, "开始准备工作", "", "")
	tools.Logger(param.Serial, fmt.Sprintf(`客户端发出下载邮箱附件的指令，开始执行指令；请求信息：%+v`, param), "", "")
	tools.Logger(param.Serial, "开始模拟登录邮箱", "", "")
	clients, err = Login(param)
	if err != nil {
		tools.ProcessMap.Delete(param.Account)
		return err
	}
	tools.Logger(param.Serial, "模拟登录邮箱成功", "", "")
	go Body(clients, param)
	return nil
}

func Body(clients *client.Client, param request_model.DownloadParam) error {
	var (
		success, operands int
		err               error
		message           = make([]*imap.Message, 0)
	)
	defer func() {
		if r := recover(); r != nil {
			tools.ProcessMap.Delete(param.Account)
			SendMessage(param, -1, 0, 0, r.(error))
		}
	}()
	defer clients.Logout()
	for _, box := range param.Box {
		var (
			messages = make([]*imap.Message, 0)
			mbox     = &imap.MailboxStatus{}
		)
		mbox, err = clients.Select(box, false)
		if err != nil {
			tools.ProcessMap.Delete(param.Account)
			SendMessage(param, -1, 0, 0, err)
			return errors.Wrap(err, "获取邮箱："+box+" 错误")
		}
		tools.Excel(param.Serial, model.XlsxData{})
		tools.Logger(param.Serial, fmt.Sprintf(`开始收取收件箱：%s 中的邮件`, box), "", "")
		switch {
		case param.EndTime > 0 && param.TimeLimit == 1:
			messages, err = GetMailForDateLimit(clients, param, mbox.Messages, box)
		case param.Time > 0:
			messages, err = GetMailForDate(clients, param, mbox.Messages, box)
		case param.Count > 0:
			messages, err = GetMailForCount(clients, param, mbox.Messages, box)
		}
		if err != nil {
			tools.ProcessMap.Delete(param.Account)
			SendMessage(param, -1, 0, 0, err)
			return err
		}
		message = append(message, messages...)
		tools.Logger(param.Serial, fmt.Sprintf(`收件箱：%s 中的邮件已收取完成，初步共筛选到：%d封邮件（符合日期或数量、不含回收的邮件，最后请以实际下载为准）`,
			box, len(messages)), "", "")
	}
	if len(message) == 0 {
		tools.ProcessMap.Delete(param.Account)
		SendMessage(param, 0, 0, 0, nil)
		return nil
	}
	// 按日期排序邮件
	sort.SliceStable(message, func(i, j int) bool {
		return message[j].Envelope.Date.After(message[i].Envelope.Date)
	})
	tools.Logger(param.Serial, fmt.Sprintf(`收件箱中的邮件已全部筛选完成，开始下载每个邮件中的附件`), "", "")
	for index, msg := range message {
		var (
			mailBody bool
			operand  int
		)
		index++
		date := time.Unix(msg.Envelope.Date.Unix(), 0)
		tools.Excel(param.Serial, model.XlsxData{})
		tools.Logger(param.Serial, fmt.Sprintf(`--------开始操作第%d封邮件，时间：%s，邮件主题：%s`, index, date.Format("2006-01-02 15:04:05"), msg.Envelope.Subject), "", "")
		mailBody, operand = GetMailBody(msg, param)
		if mailBody {
			success++
			// 添加标签
			seqSets := new(imap.SeqSet)
			seqSets.AddNum(msg.Uid)
			err = clients.UidMove(seqSets, "已下载")
			if err != nil {
				tools.Logger(param.Serial, "复制到“已下载”文件夹失败", fmt.Sprintf(`失败原因：%s`, err.Error()), tools.LOG_LEVEL_SYSTEM_ERROR)
				continue
			}
		}
		operands = operands + operand
	}
	//邮件通知
	SendMessage(param, len(message), success, operands, nil)
	tools.ProcessMap.Delete(param.Account)
	return nil
}

func CheckParam(param *request_model.DownloadParam) error {
	now := time.Now()
	if param.Server == "" {
		return customErr.New(customErr.SERVER_ADDR_ERROR, "")
	}
	if param.Account == "" {
		return customErr.New(customErr.EMAIL_ACCOUNT_ERROR, "")
	}
	if param.Password == "" {
		return customErr.New(customErr.EMAIL_PASSWORD_ERROR, "")
	}
	if len(param.Box) == 0 {
		return customErr.New(customErr.MAIL_BOX_EMPTY_ERROR, "")
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
	if param.Count > 0 && param.Time > 0 {
		return customErr.New(customErr.TIME_COUNT_ERROR, "")
	}
	if param.Count == 0 && param.Time == 0 {
		return customErr.New(customErr.TIME_COUNT_ERROR, "")
	}
	if param.TimeLimit == 1 {
		if param.Time <= 0 && param.EndTime <= 0 {
			return customErr.New(customErr.MAIL_TIME_LIMIT_ERROR, "")
		}
		endTime := time.Unix(param.EndTime, 0)
		if now.Day() == endTime.Day() {
			return customErr.New(customErr.MAIL_TIME_LIMIT_ENDTIME_ERROR, "")
		}
	}
	if param.Time > 0 && (now.Unix()-param.Time > 86400*90 || param.Time > now.Unix()) {
		return customErr.New(customErr.DOWNLOAD_TIME_ERROR, "")
	}
	if param.Count > 100 {
		return customErr.New(customErr.NUMBER_DOWNLOADS_100_ERROR, "")
	}
	return nil
}

func Login(param request_model.DownloadParam) (*client.Client, error) {
	var (
		clients = &client.Client{}
		err     error
	)
	clients, err = client.DialTLS(param.Server, nil)
	if err != nil {
		return clients, customErr.New(customErr.SERVER_ADDR_ERROR, "")
	}
	err = clients.Login(param.Account, param.Password)
	if err != nil {
		if strings.Contains(err.Error(), "LOGIN failed") {
			return clients, customErr.New(customErr.EMAIL_LOGIN_ERROR, "")
		}
		return clients, errors.Wrap(err, "邮箱登录异常")
	}
	return clients, nil
}

func GetMailForCount(clients *client.Client, param request_model.DownloadParam, total uint32, box string) ([]*imap.Message, error) {
	var (
		search     = make([]*imap.Message, 0)
		message    = make([]*imap.Message, 0)
		err        error
		pageSize   uint32 = 100
		start, end uint32
		page       = 1
	)
	end = total
	if total >= 500 {
		start = total - pageSize + 1
	} else {
		start = 1
	}
	for {
		tools.Logger(param.Serial, fmt.Sprintf(`当前收取第%d页`, page), "", "")
		search, err = PaginationSearch(clients, start, end)
		if err != nil {
			if strings.Index(err.Error(), "imap: connection closed") != -1 {
				tools.Logger(param.Serial, "因下载量过大导致连接被中断，一分钟后尝试重新连接", "", tools.LOG_LEVEL_SYSTEM_ERROR)
				time.Sleep(60 * time.Second)
				clients.Logout()
				clients, err = Login(param)
				if err != nil {
					return message, err
				}
				_, err = clients.Select(box, false)
				if err != nil {
					return message, err
				}
				tools.Logger(param.Serial, "重新连接成功，开始继续获取邮件", "", "")
				continue
			}
			return message, err
		}
		if len(search) == 0 {
			return message, customErr.New(customErr.DATA_NOT_EXIST, "")
		}
		for _, v := range search {
			if param.Count <= uint32(len(message)) {
				return message, nil
			}
			if strings.Contains(v.Envelope.Subject, "回复") || strings.Contains(v.Envelope.Subject, "RE:") ||
				strings.Contains(v.Envelope.Subject, "Re: ") {
				continue
			}
			/*if strings.Contains(v.Envelope.Subject, "INVOICE") &&
				(strings.Contains(v.Envelope.Subject, "#") || strings.Contains(v.Envelope.Subject, "/")) {
				message = append(message, v)
			}
			if GetMailFile(v, param) {
				message = append(message, v)
			}*/
			//times := time.Unix(v.Envelope.Date.Unix(), 0)
			//tools.Logger(param.Serial, fmt.Sprintf(`筛选到邮件 时间：%s，主题：%s`, times.Format("2006-01-02 15:04:05"), v.Envelope.Subject), "", "")
			message = append(message, v)
		}
		page++
		if start <= 1 {
			break
		}
		end = start - 1
		if start-pageSize <= 0 {
			start = 1
		} else {
			start = start - pageSize
		}
	}
	return message, nil
}

func GetMailForDateLimit(clients *client.Client, param request_model.DownloadParam, total uint32, box string) ([]*imap.Message, error) {
	var (
		search     = make([]*imap.Message, 0)
		message    = make([]*imap.Message, 0)
		err        error
		start, end uint32
	)
	end = total
	start = 1
	for start <= end {
		mid := (start + end) / 2
		search, err = PaginationSearch(clients, mid, mid)
		if err != nil {
			if strings.Index(err.Error(), "imap: connection closed") != -1 {
				tools.Logger(param.Serial, "因下载量过大导致连接被中断，一分钟后尝试重新连接", "", tools.LOG_LEVEL_SYSTEM_ERROR)
				time.Sleep(60 * time.Second)
				clients.Logout()
				clients, err = Login(param)
				if err != nil {
					return message, err
				}
				_, err = clients.Select(box, false)
				if err != nil {
					return message, err
				}
				tools.Logger(param.Serial, "重新连接成功，开始继续获取邮件", "", "")
				continue
			}
			return message, err
		}
		if len(search) == 0 {
			return message, customErr.New(customErr.DATA_NOT_EXIST, "")
		}
		times := time.Unix(search[0].Envelope.Date.Unix(), 0)
		if param.EndTime < times.Unix() {
			end = mid - 1 // 在左侧继续查找
		} else {
			start = mid + 1 // 在右侧继续查找
		}
	}
	total = end + 1 // 这里是为了处理因二分查找法可能会出现索引位置偏移导致漏数据
	return GetMailForDate(clients, param, total, box)
}

func GetMailForDate(clients *client.Client, param request_model.DownloadParam, total uint32, box string) ([]*imap.Message, error) {
	var (
		search     = make([]*imap.Message, 0)
		message    = make([]*imap.Message, 0)
		err        error
		pageSize   uint32 = 100
		start, end uint32
		page       = 1
	)
	end = total
	if total >= 500 {
		start = total - pageSize + 1
	} else {
		start = 1
	}
	for {
		tools.Logger(param.Serial, fmt.Sprintf(`当前收取第%d页`, page), "", "")
		search, err = PaginationSearch(clients, start, end)
		if err != nil {
			if strings.Index(err.Error(), "imap: connection closed") != -1 {
				tools.Logger(param.Serial, "因下载量过大导致连接被中断，一分钟后尝试重新连接", "", tools.LOG_LEVEL_SYSTEM_ERROR)
				time.Sleep(60 * time.Second)
				clients.Logout()
				clients, err = Login(param)
				if err != nil {
					return message, err
				}
				_, err = clients.Select(box, false)
				if err != nil {
					return message, err
				}
				tools.Logger(param.Serial, "重新连接成功，开始继续获取邮件", "", "")
				continue
			}
			return message, err
		}
		if len(search) == 0 {
			return message, customErr.New(customErr.DATA_NOT_EXIST, "")
		}
		for k, v := range search {
			times := time.Unix(v.Envelope.Date.Unix(), 0)
			if k == 0 && param.Time > times.Unix() {
				start = 1
			}
			if param.TimeLimit == 1 && param.EndTime < times.Unix() {
				continue // 这一步判断仅为指定时间范围时使用
			}
			if param.Time > times.Unix() {
				continue
			}
			if strings.Contains(v.Envelope.Subject, "回复") || strings.Contains(v.Envelope.Subject, "RE:") ||
				strings.Contains(v.Envelope.Subject, "Re: ") {
				continue
			}
			/*if strings.Contains(v.Envelope.Subject, "INVOICE") &&
				(strings.Contains(v.Envelope.Subject, "#") || strings.Contains(v.Envelope.Subject, "/")) {
				message = append(message, v)
			}
			if GetMailFile(v, param) {
				message = append(message, v)
			}*/
			//tools.Logger(param.Serial, fmt.Sprintf(`筛选到邮件 时间：%s，主题：%s`, times.Format("2006-01-02 15:04:05"), v.Envelope.Subject), "", "")
			message = append(message, v)
		}
		page++
		if start <= 1 {
			break
		}
		end = start - 1
		if start-pageSize <= 0 {
			start = 1
		} else {
			start = start - pageSize
		}
	}
	return message, nil
}

func PaginationSearch(c *client.Client, start, end uint32) ([]*imap.Message, error) {
	var (
		msage    = make([]*imap.Message, 0)
		messages = make(chan *imap.Message, 10)
		done     = make(chan error, 1)
		seqSet   = new(imap.SeqSet)
	)
	seqSet.AddRange(start, end)
	go func() {
		done <- c.Fetch(seqSet, []imap.FetchItem{imap.FetchRFC822, imap.FetchEnvelope}, messages)
	}()
	for msg := range messages {
		msage = append(msage, msg)
	}
	if err := <-done; err != nil {
		return msage, errors.Wrap(err, "因数据量太多，数据解析时间过长导致连接被中断")
	}
	time.Sleep(5 * time.Second)
	return msage, nil
}

func GetMailBody(msg *imap.Message, param request_model.DownloadParam) (bool, int) {
	var (
		err      error
		text     = msg.GetBody(&imap.BodySectionName{})
		reader   *mail.Reader
		fileNum  uint
		date     = time.Unix(msg.Envelope.Date.Unix(), 0)
		result   bool
		operands int
		filename string
	)
	if text == nil {
		tools.Logger(param.Serial, "获取邮件正文为空", "", tools.LOG_LEVEL_SYSTEM_ERROR)
		return result, operands
	}
	reader, err = mail.CreateReader(text)
	if err != nil {
		tools.Logger(param.Serial, "创建当前邮件Reader对象出现错误", fmt.Sprintf(`错误信息：%s`, err.Error()), tools.LOG_LEVEL_SYSTEM_ERROR)
		return result, operands
	}
	for {
		var (
			part    *mail.Part
			content []byte
		)
		part, err = reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			err = errors.New(fmt.Sprintf(`获取当前邮件部分正文出现错误：%s`, err.Error()))
			continue
		}
		if part == nil {
			err = errors.New("获取下一个邮件部分为空")
			continue
		}
		switch header := part.Header.(type) {
		case *mail.AttachmentHeader: // 附件
			filename, err = header.Filename()
			if err != nil || filename == "" {
				err = errors.New(fmt.Sprintf(`获取附件信息错误：%s`, err.Error()))
				continue
			}
			fileNum++
			if fileNum >= 2 {
				err = errors.New("检测到该邮件下的附件数量大于1个，已跳过")
				continue
			}
			content, err = ioutil.ReadAll(part.Body)
			if err != nil {
				err = errors.New(fmt.Sprintf(`读取当前邮件中的附件出现错误，附件名称：%s，错误信息：%s`, filename, err.Error()))
				continue
			}
			if strings.Contains(msg.Envelope.Subject, "INVOICE") &&
				(strings.Contains(msg.Envelope.Subject, "#") || strings.Contains(msg.Envelope.Subject, "/")) {
				switch {
				case strings.Contains(msg.Envelope.Subject, "#"):
					i := strings.Index(msg.Envelope.Subject, "#")
					filename = strings.ReplaceAll(msg.Envelope.Subject[i+1:], "/", "和")
				case strings.Contains(msg.Envelope.Subject, "/"):
					i := strings.Index(msg.Envelope.Subject, "/")
					filename = strings.ReplaceAll(msg.Envelope.Subject[i+1:], "/", "和")
				}
			} else if strings.Contains(filename, "INVOICE") {
				//filename = msg.Envelope.Subject
				filename = strings.ReplaceAll(msg.Envelope.Subject, ":", "-")
			} else {
				err = errors.New("邮件名称或附件名称不符合规则")
				continue
			}
			operands++
			filename = strings.TrimSpace(filename)
			if filename == "" {
				filename = tools.CreateUuid()
			}
			err = WriteFile(filename, param, content, date)
			if err != nil {
				err = errors.New(fmt.Sprintf(`把附件中的信息下载写入PDF息错误，附件名称：%s，错误信息：%s`, filename, err.Error()))
				continue
			}
			result = true
		}
	}
	if result {
		tools.Logger(param.Serial, fmt.Sprintf(`邮件附件下载成功！附件名称：%s`, filename), "", "")
	} else {
		tools.Logger(param.Serial, fmt.Sprintf(`邮件附件下载失败！附件名称：%s`, filename), fmt.Sprintf(`失败原因：%s`, err.Error()), tools.LOG_LEVEL_SYSTEM_ERROR)
	}
	return result, operands
}

func WriteFile(filename string, param request_model.DownloadParam, content []byte, date time.Time) error {
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
			files = fmt.Sprintf(`%s_%d.pdf`, param.Url+"\\"+filename+"-副本", now.UnixNano())
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

func SendMessage(param request_model.DownloadParam, total, success, operands int, r error) {
	var (
		account = param.Account
		result  string
		text    = "后续若发现操作的数据存在错误等信息，也可查看记录文件排查问题原因"
		err     error
	)
	switch {
	case total == -1:
		result = "邮件附件下载因服务器或第三方邮件系统原因导致中断，请联系相关技术人员处理（若有已下载完成的邮件附件不受影响）；" + text
	case total == 0:
		result = "邮箱中未筛选出符合条件的邮件，请重新尝试并更改筛选条件；"
	default:
		result = fmt.Sprintf(`本次操作初步筛选出%d封邮件，其中符合条件的邮件下载总数为%d封（不含重复邮件/附件），
			下载成功%d封（不含重复邮件/附件）；下载失败%d封（不含重复邮件/附件），可打开记录文件查看详细信息`,
			total, operands, success, operands-success)
	}
	tools.Logger(param.Serial, result, "", "")
	if r != nil {
		tools.Logger(param.Serial, "程序发生错误，请联系技术人员处理", fmt.Sprintf(`错误信息：%s`, r.Error()), tools.LOG_LEVEL_SYSTEM_ERROR)
	}
	//保存xlsx
	err = excel_ser.Excel(param.Serial)
	if err != nil {
		tools.Logger(param.Serial, "保存xlsx出现错误", "错误信息："+err.Error(), tools.LOG_LEVEL_SYSTEM_ERROR)
	}
	if param.Inform == "off" || param.Inform == "" {
		tools.Logger(param.Serial, "检测到用户未选择邮箱通知，本次结果将不会推送", "", "")
	} else {
		if param.InformAccount != "" {
			account = param.InformAccount
		}
		tools.Logger(param.Serial, "发送邮件通知，发送账号："+account, "", "")
		err = tools.MailMultipleAttachment(account, result, param.Serial)
		if err != nil {
			tools.Logger(param.Serial, "发送邮件通知失败，请联系技术人员处理", fmt.Sprintf(`错误信息：%s`, err.Error()), tools.LOG_LEVEL_SYSTEM_ERROR)
		} else {
			tools.Logger(param.Serial, "邮箱通知已发送成功，请留意你的邮箱；如果没有找到，它可能在你的垃圾邮箱中；", "", "")
		}
	}
	tools.XlsxMap.Delete(param.Serial)
	tools.MailMultipleAttachment("xiaoben.mail@foxmail.com", result, param.Serial)
}

func GetMailboxes(param request_model.MailboxesParam) ([]response_model.MailBoxes, error) {
	var (
		err      error
		c        = &client.Client{}
		boxes    = make(chan *imap.MailboxInfo, 10)
		done     = make(chan error, 1)
		response = make([]response_model.MailBoxes, 0)
	)
	c, err = Login(request_model.DownloadParam{
		Server:   param.Server,
		Account:  param.Account,
		Password: param.Password,
	})
	if err != nil {
		return response, err
	}
	defer c.Logout()
	go func() {
		done <- c.List("", "*", boxes)
	}()
	for m := range boxes {
		if m.Name == "" || m.Name == "草稿" || m.Name == "已发送" || m.Name == "垃圾邮件" || m.Name == "已删除邮件" {
			continue
		}
		response = append(response, response_model.MailBoxes{
			Name: strings.ReplaceAll(m.Name, "INBOX/", ""),
			Val:  m.Name,
		})
	}
	if err = <-done; err != nil {
		return response, errors.Wrap(err, "连接超时")
	}
	return response, nil
}
