package main

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/pkg/errors"
	"mail_download/model/request_model"
	"sort"
	"time"
)

func GetEmailByUid(param request_model.DownloadParam) error {
	var (
		c          = &client.Client{}
		err        error
		mbox       = &imap.MailboxStatus{}
		start, end uint32
		search     = make([]*imap.Message, 0)
		message    = make([]*imap.Message, 0)
		pageSize   uint32
	)
	// 连接到邮箱的 IMAP 服务器
	c, err = client.DialTLS(param.Server, nil)
	if err != nil {
		return fmt.Errorf("Error connecting to server: %v", err)
	}
	defer c.Logout()
	if err = c.Login(param.Account, param.Password); err != nil {
		return fmt.Errorf("Error logging in: %v", err)
	}
	fmt.Println("-------登录成功")
	// 选择收件箱
	mbox, err = c.Select("INBOX", false)
	if err != nil {
		return fmt.Errorf("Error selecting mailbox: %v", err)
	}
	//total = mbox.Messages
	end = mbox.Messages
	pageSize = 100
	for {
		start = mbox.Messages - pageSize + 1
		search, err = UidSearch(c, start, end)
		if err != nil {
			return err
		}
		if len(search) == 0 {
			return fmt.Errorf("获取邮件信息出现问题")
		}
		//检查第一条的数据：如果第一条的数据小于了客户端请求的数据，说明数据就在查询到的着后面一半的数据之中
		if param.Time > search[0].Envelope.Date.Unix() {
			for _, v := range search {
				times := time.Unix(v.Envelope.Date.Unix(), 0)
				fmt.Println(v.Envelope.Date.Format("2006-01-02 15:04:05"))
				fmt.Println(times.Format("2006-01-02 15:04:05"))
				fmt.Println("========")
				if param.Time <= times.Unix() {
					message = append(message, v)
				}
			}
			break
		}
		message = append(message, search...)
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
		fmt.Println("----------------------------------")
	}
	fmt.Println(fmt.Sprintf(`共筛选到：%d封邮件；`, len(message)))

	return nil
}

func UidSearch(c *client.Client, start, end uint32) ([]*imap.Message, error) {
	var (
		msage    = make([]*imap.Message, 0)
		messages = make(chan *imap.Message, 10)
		done     = make(chan error, 1)
		seqSet   = new(imap.SeqSet)
	)
	for uid := start; uid <= end; uid++ {
		seqSet.AddNum(uid)
	}
	go func() {
		done <- c.UidFetch(seqSet, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()
	for msg := range messages {
		//GetMailFile(msg)
		text := msg.GetBody(&imap.BodySectionName{})
		if text == nil {
			fmt.Println("获取邮件正文为空")
			continue
		}
		msage = append(msage, msg)
	}
	if err := <-done; err != nil {
		return msage, errors.Wrap(err, "因数据量太多，数据解析时间过长导致连接被中断")
	}
	return msage, nil
}
