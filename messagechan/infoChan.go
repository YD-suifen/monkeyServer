package messagechan

import "monkeyServer/dataTypeStruck"

var AlarmInfoChan chan []dataTypeStruck.AlarmInfo

func init()  {

	AlarmInfoChan = make(chan []dataTypeStruck.AlarmInfo,100)

}