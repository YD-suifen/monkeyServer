package alarm

import (
	"monkeyServer/dataTypeStruck"
	"monkeyServer/messagechan"
)

func ReadAlarmInfo() (bool, dataTypeStruck.AlarmInfo) {
	data, ok  := <- messagechan.AlarmInfoChan
	if ok {
		return true,data
	}
	return false,data
}