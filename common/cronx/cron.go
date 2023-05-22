package cronx

import (
	"github.com/robfig/cron"
)

var Cron *cron.Cron

func InitCron() {
	Cron = cron.New()
}
