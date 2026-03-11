package main

import (
	"github.com/robfig/cron/v3"
)

type Job struct {
	Scheduled string
	Func      func()
}

func MidnightJob() {
	SweepSessions()
}

func CronJobs() {
	c := cron.New()
	c.AddFunc("0 0 0 * * *", MidnightJob)
	c.Start()
}
