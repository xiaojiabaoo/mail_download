package model

type AppConfig struct {
	Name     string `json:"name"`
	Port     string `json:"port"`
	CronTime string `json:"cron_time"`
}
