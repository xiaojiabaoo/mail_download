package main

import (
	"mail_download/common/configx"
	"mail_download/router"
)

func main() {
	//初始化Viper，获取到配置文件路径和名字
	configx.InitViper()
	//初始化配置数据
	configx.InitConfigData()
	//初始化定时器
	//cronx.InitCron()
	//初始化Job服务
	//job.InitJob()
	//服务启动端口
	router.Run("0.0.0.0:" + "8023")
}
