package dao

import (
	"encoding/json"
	"fmt"
	"monkeyServer/dataTypeStruck"
	"monkeyServer/logUtils"
	"monkeyServer/utils"
)

func AlarmSelect(times int64) [][]byte {

	var data [][]byte

	cpu := CpuAlarm(times)
	mem := MemAlarm(times)
	tcp := TcpnetAlarm(times)
	data = append(data,cpu,mem,tcp)
	//dataRespne = append(append(cpu,mem...),tcp...)
	//dataJson, _ := json.Marshal(dataRespne)
	return data
}

func CpuAlarm(times int64) []byte {



	db := utils.SqlxCli()
	defer db.Close()

	var dataList []dataTypeStruck.NoAlarmCpu
	sql := fmt.Sprintf("select hostName,keyName,usedCpu,timeUnix from monkey_s_cpudata where timeUnix = %v",times)
	logUtils.Debugf("cpuAlarm sql=%v",sql)

	if err := db.Select(&dataList,sql);err != nil{
		logUtils.Errorf("cpuAlarm Time=%v,error=%v",times,err)
	}
	var data dataTypeStruck.NoAlarmInfo
	data.Type = 1
	data.Data = dataList
	jsonData, _ := json.Marshal(data)
	return jsonData
}

func MemAlarm(times int64) []byte {

	db := utils.SqlxCli()
	defer db.Close()

	var dataList []dataTypeStruck.NoAlarmMem
	sql := fmt.Sprintf("select hostName,keyName,used,timeUnix from monkey_s_memdata where timeUnix = %v",times)
	logUtils.Debugf("memAlarm sql=%v",sql)

	if err := db.Select(&dataList,sql);err != nil{
		logUtils.Errorf("memAlarm Time=%v,error=%v",times,err)
	}
	var data dataTypeStruck.NoAlarmInfo
	data.Type = 2
	data.Data = dataList
	jsonData, _ := json.Marshal(data)
	return jsonData
}
//
//func diskAlarm(times int64) []byte {
//
//	db := utils.SqlxCli()
//	defer db.Close()
//
//	var dataList []dataTypeStruck.DiskRespone
//	sql := fmt.Sprintf("select hostName,used,timeUnix from monkey_s_memdata where timeUnix = %v",times)
//	logUtils.Debugf("memAlarm sql=%v",sql)
//
//	if err := db.Select(&dataList,sql);err != nil{
//		logUtils.Errorf("memAlarm Time=%v,error=%v",times,err)
//		return nil
//	}
//	jsonData, _ := json.Marshal(dataList)
//	return jsonData
//}



func TcpnetAlarm(times int64) []byte {

	db := utils.SqlxCli()
	defer db.Close()

	var dataList []dataTypeStruck.NoAlarmTcp
	sql := fmt.Sprintf("select hostName,keyName,allConn,timeUnix from monkey_s_tcpnetdata where timeUnix = %v",times)
	logUtils.Debugf("tcpnetAlarm sql=%v",sql)

	if err := db.Select(&dataList,sql);err != nil{
		logUtils.Errorf("tcpnetAlarm Time=%v,error=%v",times,err)
		return nil
	}
	var data dataTypeStruck.NoAlarmInfo
	data.Type = 3
	data.Data = dataList
	jsonData, _ := json.Marshal(data)
	return jsonData
}