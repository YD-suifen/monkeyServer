package alarm

import (
	"encoding/json"
	"fmt"
	"monkeyServer/dao"
	"monkeyServer/dataTypeStruck"
	"monkeyServer/messagechan"
	"monkeyServer/utils"
	"time"
)

type SData dataTypeStruck.NoAlarmInfo
type DData dataTypeStruck.AlarmInfo
//type ByteS []byte

//func init()  {
//	go AlarmActive()
//}

func AlarmActive()  {

	time.Sleep(time.Second * 30)
	for  {
		fmt.Println("check AlarmActive")
		time.Sleep(time.Second * 60)
		times := utils.BeMin()
		byteListData := dao.AlarmSelect(times)
		for _, v := range byteListData{
			var SDatas SData
			SDatas = get(v)
			SDatas.Get()
		}
	}

}

func get(data []byte) SData {
	var sData SData
	_ = json.Unmarshal(data,&sData)
	return sData
}


func (c *SData) Get() {
	switch c.Type {
	case 1:
		var sCpu []dataTypeStruck.NoAlarmCpu
		var Odata []dataTypeStruck.AlarmInfo
		jsonData,_ := json.Marshal(c.Data)
		_ = json.Unmarshal(jsonData,&sCpu)
		for _, v := range sCpu{
			if alarmDataActive(v.Used) {
				continue
			}
			Ddata := dataTypeStruck.AlarmInfo{}
			Ddata.KeyName = v.KeyName
			Ddata.HostName = v.HostName
			Ddata.Type = c.Type
			Ddata.State = false
			Ddata.Used = v.Used
			Ddata.TimeUnix = v.TimeUnix
			Odata = append(Odata,Ddata)
		}
		messagechan.AlarmInfoChan <- Odata
		break

	case 2:
		var sCpu []dataTypeStruck.NoAlarmMem
		var Odata []dataTypeStruck.AlarmInfo
		jsonData,_ := json.Marshal(c.Data)
		_ = json.Unmarshal(jsonData,&sCpu)
		for _, v := range sCpu{
			if alarmDataActive(v.Used) {
				continue
			}
			Ddata := dataTypeStruck.AlarmInfo{}
			Ddata.HostName = v.HostName
			Ddata.KeyName = v.KeyName
			Ddata.Type = c.Type
			Ddata.State = false
			Ddata.Used = v.Used
			Ddata.TimeUnix = v.TimeUnix
			Odata = append(Odata,Ddata)

		}
		messagechan.AlarmInfoChan <- Odata
		break
	case 3:

	}

}

func alarmDataActive(value float64) bool {
	if value > 60  || value <= 0 {
		return false
	}
	return true
}