package job

import (
	"mail_download/common/cronx"
)

func InitJob() {
	initCron()
}

func initCron() {
	/*err := cronx.Cron.AddFunc(configx.AppConfigData.CronTime, message_ser.Push)
	if err != nil {
		fmt.Println("定时器错误：", err)
	}*/
	cronx.Cron.Start()
}
