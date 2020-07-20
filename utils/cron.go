package utils

import (
	"github.com/robfig/cron/v3"
	"monkeyServer/logUtils"
	"monkeyServer/server/Calculation"
)


var Crontab *cron.Cron



func init()  {
	logUtils.Info("init Crontab")

	Crontab = cron.New(cron.WithSeconds())
	Crontab.Start()

}

func AddTasks() {

	spec := "0 0 12 * * ?"
	_, err := Crontab.AddFunc(spec, func() {
		Calculation.Algorithm()
	})
	if err != nil{
		logUtils.Errorf("AddTasks error=%v",err)
	}
}