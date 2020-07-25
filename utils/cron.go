package utils

import (
	"github.com/robfig/cron/v3"
	"monkeyServer/logUtils"
)


//var Crontab *cron.Cron



func TasksInit() *cron.Cron {
	logUtils.Info("init Crontab")

	Crontab := cron.New(cron.WithSeconds())

	return Crontab
}

//func AddTasks()  {
//	spec := "0 40 17 * * ?"
//	_, err := Crontab.AddFunc(spec, func() {
//		Calculation.Algorithm()
//	})
//	if err != nil{
//		logUtils.Errorf("AddTasks error=%v",err)
//	}
//
//}