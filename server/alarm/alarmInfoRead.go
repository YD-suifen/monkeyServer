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
	case data := <- messagechan.AlarmInfoChan:
		fmt.Println("read over")
		if data != nil {
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
//
//func ReadAlarmInfo() (bool, []dataTypeStruck.AlarmInfo) {
//	fmt.Println("read start")
//	logUtils.Info("read alarmInfo")
//
//	var data []dataTypeStruck.AlarmInfo
//	var data2 dataTypeStruck.AlarmInfo
//	var data3 dataTypeStruck.AlarmInfo
//
//	data2.HostName = "jinage"
//	data2.KeyName = "songyuan"
//	data2.TimeUnix = 8677888
//	data2.Type = 1
//	data2.Used = 29.0
//	data2.State = true
//
//	data3.HostName = "jiange2"
//	data3.KeyName = "chagnchun"
//	data3.TimeUnix = 8677888
//	data3.Type = 1
//	data3.Used = 29.0
//	data3.State = true
//	data = append(data,data2,data3)
//	fmt.Println(len(data))
//	return true,data
//
//}