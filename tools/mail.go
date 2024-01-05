package tools

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

// SendMail 发送邮箱消息
// user 发送账号
// body 发送内容
// subject 发送邮箱主题
func SendMail(account []string, body, subject string) error {
	var (
		user                        = "xiaoben_mail@163.com"
		password                    = "TRCTOKIPVNJRDXOF"
		host                        = "smtp.163.com:25"
		contentType, replyToAddress string
		cc, bcc                     []string
	)
	if len(account) == 0 || len(body) == 0 || len(subject) == 0 {
		return errors.New("发送信息不能为空")
	}
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	ccAddress := strings.Join(cc, ";")
	bccAddress := strings.Join(bcc, ";")
	toAddress := strings.Join(account, ";")
	msg := []byte("To: " + toAddress + "\r\nFrom: " + "CCL邮件操作<" + user + ">" + "\r\nSubject: " + subject + "\r\nReply-To: " + replyToAddress + "\r\nCc: " + ccAddress + "\r\nBcc: " + bccAddress + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := MergeSlice(account, cc)
	sendTo = MergeSlice(sendTo, bcc)
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	if err != nil {
		return errors.Wrap(err, "邮箱发送失败！请检查邮箱是否填写正确无误")
	}
	return err
}

func MergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}

func MailAttachment(email, body, serial string) error {
	var (
		subject string
		path    = fmt.Sprintf(`./logs/%s.log`, serial)
	)
	if strings.HasPrefix(serial, "CCL") {
		subject = fmt.Sprintf(`CCL流程操作结果通知，流水号：%s`, serial)
	} else {
		subject = fmt.Sprintf(`CCL流程邮件下载结果通知，流水号：%s`, serial)
	}
	return SendMailAttachment([]string{email}, subject, body, path)
}

func SendMailAttachment(account []string, subject, body, path string) error {
	var (
		user     = "xiaoben_mail@163.com"
		password = "TRCTOKIPVNJRDXOF"
		host     = "smtp.163.com:25"
		err      error
		message  = bytes.NewBuffer(nil) // 创建邮件消息体
		name     = filepath.Base(path)
		file     []byte
		accounts string
		cc, bcc  []string
	)
	// 读取附件文件
	file, err = os.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "读取文件错误")
	}
	accounts = strings.Join(account, ";")
	// 创建附件消息体
	part := fmt.Sprintf(`Content-Disposition: attachment; filename="%s"`, name)
	message.WriteString(fmt.Sprintf("To: %s\r\n", accounts))
	message.WriteString(fmt.Sprintf("From: %s\r\n", "CCL邮件操作<"+user+">"))
	message.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	message.WriteString("MIME-version: 1.0\r\n")
	message.WriteString(fmt.Sprintf("Content-Type: multipart/mixed;charset=UTF-8; boundary=\"%s\"\r\n", "boundary_separator"))
	message.WriteString("\r\n")
	message.WriteString(fmt.Sprintf("--%s\r\n", "boundary_separator"))
	message.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	message.WriteString("\r\n")
	// 添加邮件正文
	message.WriteString(body)
	// 添加附件内容
	message.WriteString(fmt.Sprintf("\r\n--%s\r\n", "boundary_separator"))
	message.WriteString(fmt.Sprintf("%s\r\n", part))
	message.WriteString(fmt.Sprintf("Content-Type: application/octet-stream\r\n"))
	message.WriteString(fmt.Sprintf("Content-Transfer-Encoding: base64\r\n"))
	message.WriteString("\r\n")
	// 添加附件数据
	encoded := base64.StdEncoding.EncodeToString(file)
	message.WriteString(encoded)
	message.WriteString("\r\n")
	message.WriteString(fmt.Sprintf("\r\n--%s--\r\n", "boundary_separator"))
	// 连接 SMTP 服务器并发送附件
	auth := smtp.PlainAuth("", user, password, strings.Split(host, ":")[0])
	sendTo := MergeSlice(MergeSlice(account, cc), bcc)
	err = smtp.SendMail(host, auth, user, sendTo, message.Bytes())
	if err != nil {
		return errors.Wrap(err, "邮箱发送失败！请检查邮箱是否填写正确无误")
	}
	return nil
}
