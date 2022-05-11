package cron

import (
	"counter/V2"
	"github.com/robfig/cron"
)

func InitCron() *cron.Cron {
	c := cron.New()
	//week周更新 周日零点更新
	c.AddFunc("0 0 0 ? * SUN", func() {
		V2.Cts.ResetIndex("week")
	})
	//month月更新,每月一号更新
	c.AddFunc("0 0 0 1 * ? ", func() {
		V2.Cts.ResetIndex("month")

	})

	return c
}
