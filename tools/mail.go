package tools

import (
	"github.com/pkg/errors"
	"net/smtp"
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
	//subject := "邮件验证码：" + codeNum
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	ccAddress := strings.Join(cc, ";")
	bccAddress := strings.Join(bcc, ";")
	toAddress := strings.Join(account, ";")
	subject = "附件下载结果：" + subject
	msg := []byte("To: " + toAddress + "\r\nFrom: " + "邮件附件下载<" + user + ">" + "\r\nSubject: " + subject + "\r\nReply-To: " + replyToAddress + "\r\nCc: " + ccAddress + "\r\nBcc: " + bccAddress + "\r\n" + contentType + "\r\n\r\n" + body)
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
