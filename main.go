package main

import (
	"flag"
	"fmt"
	"monkeyServer/logUtils"
	"monkeyServer/router"
	"monkeyServer/utils"
	"monkeyServer/server/Task"
	"monkeyServer/server/alarm"
	"os"
)

var (
	H bool
	C string
)
func init() {
	flag.BoolVar(&H, "h", false, "help")
	flag.StringVar(&C, "c", "", "read file path!")
	flag.Usage = usage
}
func usage() {
	fmt.Fprintf(os.Stderr, `
Usage: go-execl [-h] [-c config] [-o savedir]

Options:
`)
	flag.PrintDefaults()
}
func main() {

	logUtils.InitLogger()
	logUtils.SugarLogger.Info("log init ok",)

	flag.Parse()
	if H {
		flag.Usage()
		return
	}
	utils.InitRedisConfigs(C)
	logUtils.SugarLogger.Info("config init ok")
	Task.Task()
	go alarm.AlarmActive()

	fmt.Println("aaa",utils.Config.Master,utils.Config.Port,utils.Config.DB.DbHost,utils.Config.Tasks.Trend)
	logUtils.SugarLogger.Info("host ColleData start")
	router.RegistRouter()
}