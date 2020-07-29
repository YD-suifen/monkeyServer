package alarm

import (
	"fmt"
	"monkeyServer/dataTypeStruck"
	"monkeyServer/logUtils"
	"monkeyServer/messagechan"
)

func ReadAlarmInfo() (bool, []dataTypeStruck.AlarmInfo) {
	fmt.Println("read start")
	logUtils.Info("read alarmInfo")

	select {
	case data, ok  := <- messagechan.AlarmInfoChan:
		fmt.Println("read over")
		if ok {
			fmt.Println("read true")
			return true,data
		}
		fmt.Println("read false")
		return false,nil
	default:
		fmt.Println("read nil false")
		return false, nil
	}
}

//func ReadAlarmInfo() (bool, dataTypeStruck.AlarmInfo) {
//	fmt.Println("read start")
//	logUtils.Info("read alarmInfo")
//
//	var data dataTypeStruck.AlarmInfo
//	data.HostName = "jinage"
//	data.TimeUnix = 8677888
//	data.Type = 1
//	data.Used = 29.0
//	data.State = true
//	return true,data
//
//}