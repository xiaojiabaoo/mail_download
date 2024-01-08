package main

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-message/mail"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"mail_download/model/request_model"
	"mail_download/tools"
	customErr "mail_download/tools/error"
	"os"
	"strings"
	"time"
)

func GetMailFile(msg *imap.Message, param request_model.DownloadParam) {
	var (
		err     error
		text    = msg.GetBody(&imap.BodySectionName{})
		reader  *mail.Reader
		fileNum uint
	)
	if text == nil {
		fmt.Println("获取邮件正文为空")
		return
	}
	reader, err = mail.CreateReader(text)
	if err != nil {
		fmt.Println("创建当前邮件Reader对象出现错误：" + err.Error())
		return
	}
	// 处理邮件正文
	for {
		var (
			part     *mail.Part
			filename string
			content  []byte
		)
		part, err = reader.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("获取当前邮件部分正文出现错误:" + err.Error())
			continue
		}
		if part == nil {
			fmt.Println("获取下一个邮件部分为空")
			continue
		}
		switch header := part.Header.(type) {
		case *mail.AttachmentHeader: // 附件
			filename, err = header.Filename()
			if err != nil || filename == "" {
				fmt.Println("获取当前邮件部分正文出现错误:" + err.Error())
				continue
			}
			fileNum++
			if fileNum >= 2 {
				fmt.Println("检测到该邮件下的附件数量大于1个，已跳过")
				continue
			}
			content, err = ioutil.ReadAll(part.Body)
			if err != nil {
				fmt.Println(fmt.Sprintf(`读取当前邮件中的附件出现错误，附件名称：%s，错误信息：%s`, filename, err.Error()))
				continue
			}
			if strings.Contains(msg.Envelope.Subject, "#") {
				i := strings.Index(msg.Envelope.Subject, "#")
				filename = strings.ReplaceAll(msg.Envelope.Subject[i+1:], "/", "和")
			} else if strings.Contains(msg.Envelope.Subject, "/") {
				i := strings.Index(msg.Envelope.Subject, "/")
				filename = strings.ReplaceAll(msg.Envelope.Subject[i+1:], "/", "和")
			}
			filename = strings.TrimSpace(filename)
			if filename == "" {
				filename = tools.CreateUuid()
			}
			date := time.Unix(msg.Envelope.Date.Unix(), 0)
			err = writeFile(filename, param, content, date)
			if err != nil {
				fmt.Println("发送日期：" + date.Format("2006-01-02 15:04:05") + "；邮件主题：" + msg.Envelope.Subject + "；附件名称：" + filename + "；错误信息：" + err.Error())
				continue
			}
		}
	}
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
