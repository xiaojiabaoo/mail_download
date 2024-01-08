package main

import (
	"fmt"
	"github.com/emersion/go-imap"
	goImapId "github.com/emersion/go-imap-id"
	"github.com/emersion/go-imap/client"
	"github.com/pkg/errors"
	"mail_download/model/request_model"
	"net"
	"sort"
	"strings"
	"time"
)

func GetEmailByRange(param request_model.DownloadParam) error {
	var (
		c               = &client.Client{}
		err             error
		mbox            = &imap.MailboxStatus{}
		start, end      uint32
		search          = make([]*imap.Message, 0)
		message         = make([]*imap.Message, 0)
		pageSize, total uint32
		dial            = new(net.Dialer)
	)
	// 连接到邮箱的 IMAP 服务器
	dial.Timeout = time.Duration(3) * time.Second
	c, err = client.DialWithDialerTLS(dial, param.Server, nil)
	if err != nil {
		return fmt.Errorf("Error connecting to server: %v", err)
	}
	defer c.Logout()
	if err = c.Login(param.Account, param.Password); err != nil {
		return fmt.Errorf("Error logging in: %v", err)
	}
	fmt.Println("-------登录成功")
	newClient := goImapId.NewClient(c)
	newClient.ID(
		goImapId.ID{
			goImapId.FieldName:    "IMAPClient",
			goImapId.FieldVersion: "2.1.0",
		},
	)
	// 选择收件箱
	mbox, err = c.Select("INBOX", false)
	if err != nil {
		return fmt.Errorf("Error selecting mailbox: %v", err)
	}
	total = mbox.Messages
	end = mbox.Messages
	pageSize = 100
	if total >= 500 {
		start = total - pageSize + 1
	} else {
		start = 1
	}
	for {
		search, err = RangeSearch(c, start, end)
		if err != nil {
			return err
		}
		if len(search) == 0 {
			return fmt.Errorf("获取邮件信息出现问题")
		}
		for _, v := range search {
			if !strings.Contains(v.Envelope.Subject, "INVOICE") || strings.Contains(v.Envelope.Subject, "回复") ||
				strings.Contains(v.Envelope.Subject, "RE:") || strings.Contains(v.Envelope.Subject, "Re: ") {
				continue
			}
			if !strings.Contains(v.Envelope.Subject, "#") && !strings.Contains(v.Envelope.Subject, "/") {
				continue
			}
			times := time.Unix(v.Envelope.Date.Unix(), 0)
			if param.Time <= times.Unix() {
				message = append(message, v)
			}
		}
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
	// 按日期排序邮件
	sort.SliceStable(message, func(i, j int) bool {
		return message[j].Envelope.Date.After(message[i].Envelope.Date)
	})
	for _, v := range message {
		fmt.Println("Subject:", v.Envelope.Subject)
		date := time.Unix(v.Envelope.Date.Unix(), 0)
		fmt.Println("Date:", date.Format("2006-01-02 15:04:05"))
		GetMailFile(v, param)
		fmt.Println("----------------------------------")
	}
	fmt.Println(fmt.Sprintf(`共筛选到：%d封邮件；`, len(message)))
	return nil
}

func RangeSearch(c *client.Client, start, end uint32) ([]*imap.Message, error) {
	var (
		seqSet   = new(imap.SeqSet)
		msage    = make([]*imap.Message, 0)
		messages = make(chan *imap.Message, 10)
		done     = make(chan error, 1)
	)
	seqSet.AddRange(start, end)
	go func() {
		done <- c.Fetch(seqSet, []imap.FetchItem{imap.FetchRFC822, imap.FetchEnvelope}, messages)
		//done <- c.Fetch(seqSet, []imap.FetchItem{imap.FetchBodyStructure, imap.FetchEnvelope}, messages)
	}()
	for msg := range messages {
		msage = append(msage, msg)
	}
	if err := <-done; err != nil {
		return msage, errors.Wrap(err, "因数据量太多，数据解析时间过长导致连接被中断")
	}
	return msage, nil
}
