package configx

import (
	"fmt"
	"github.com/spf13/viper"
	"mail_download/model"
	"os"
)

var (
	AppConfigData model.AppConfig
)

func InitViper() {
	confPath, _ := os.Getwd()
	viper.AddConfigPath(confPath)
	viper.SetConfigName("app")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("初始化Viper出错：", err)
	}
}

func InitConfigData() {
	//读取APP配置数据
	AppConfigData.Name = viper.GetString("app.name")
	AppConfigData.Port = viper.GetString("app.port")
	AppConfigData.CronTime = viper.GetString("app.cron_time")
}
