package main

import (
	"time"

	"github.com/robfig/cron/v3"
)

var cronOpt = cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor

func getNextTime(rule string, cronOpt cron.ParseOption) (time.Time, error) {
	specParser := cron.NewParser(cronOpt)
	sched, err := specParser.Parse(rule)
	if err != nil {
		return time.Time{}, err
	}
	t := sched.Next(time.Now())
	return t, nil
}
