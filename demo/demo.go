package main

import (
	"fmt"
	"mail_download/model/request_model"
	"time"
)

func main() {
	param := request_model.DownloadParam{
		Server:        "imap.mxhichina.com:993",
		Account:       "acc02@ggt-world.com",
		Password:      "Acc01.ggt125",
		Time:          1701360000, // 2024-01-01 00:00:00
		Count:         0,
		Url:           "C:\\mail_file_download",
		Classify:      "month",
		Type:          "cover",
		Inform:        "on",
		InformAccount: "xiaoben.mail@qq.com",
		Serial:        fmt.Sprintf(`%d`, time.Now().Unix()),
	}
	err := GetEmailByUid(param)
	if err != nil {
		fmt.Println(err)
	}
}
