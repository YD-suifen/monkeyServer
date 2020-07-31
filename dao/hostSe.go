package dao

import (
	"encoding/json"
	"fmt"
	"monkeyServer/dataTypeStruck"
	"monkeyServer/logUtils"
	"monkeyServer/utils"
)

func HostAll(hostName string,startTime,endTime int64) map[string][]byte {

	data := make(map[string][]byte)
	data["cpu"] = HSelectCpu(hostName,startTime,endTime)
	data["mem"] = HSelectMem(hostName,startTime,endTime)
	data["disk"] = HSelectDisk(hostName,startTime,endTime)
	data["tcp"] = HSelectTcpnet(hostName,startTime,endTime)
	return data
}


func HSelectCpu(hostName string,startTime,endTime int64) []byte {


	db := utils.SqlxCli()
	defer db.Close()
	var data []dataTypeStruck.CpuRespone

	sql := fmt.Sprintf("select hostName,usedCpu,timeUnix from monkey_s_cpudata where hostName = '%v' and timeUnix >= %v and timeUnix < %v",hostName,startTime,endTime)
	logUtils.Debugf("HSelectCpu sql=%v",sql)

	if err := db.Select(&data,sql);err != nil{
		logUtils.Errorf("HSelectCpu hostName=%v,startTime=%v,error=%v",hostName,startTime,err)
		return nil
	}
	jsonData, _ := json.Marshal(data)
	return jsonData
}

func HSelectMem(hostName string,startTime,endTime int64) []byte {
	db := utils.SqlxCli()
	defer db.Close()
	var data []dataTypeStruck.MemRespone

	sql := fmt.Sprintf("select hostName,used,timeUnix from monkey_s_memdata where hostName = '%v' and timeUnix >= %v and timeUnix < %v",hostName,startTime,endTime)
	if err := db.Select(&data,sql);err != nil{
		logUtils.Errorf("tvSelectMem keyName=%v,startTime=%v,error=%v",hostName,startTime,err)
		return nil
	}
	jsonData, _ := json.Marshal(data)
	return jsonData


}

func HSelectDisk(hostName string,startTime,endTime int64) []byte {
	db := utils.SqlxCli()
	defer db.Close()
	var dbData []dataTypeStruck.DiskDB
	var datas []dataTypeStruck.DiskRespone

	sql := fmt.Sprintf("select hostName,disk,timeUnix from monkey_s_diskdata where hostName = '%v' and timeUnix >= %v and timeUnix < %v",hostName,startTime,endTime)
	if err := db.Select(&dbData,sql);err != nil{
		logUtils.Errorf("tvSelectDisk keyName=%v,startTime=%v,error=%v",hostName,startTime,err)
		return nil
	}
	for _, v := range dbData {
		disk := &[]dataTypeStruck.Disk{}
		data := dataTypeStruck.DiskRespone{}
		_ = json.Unmarshal([]byte(v.Disk),disk)
		data.HostName = v.HostName
		data.TimeUnix = v.TimeUnix
		data.Disks = *disk
		datas = append(datas,data)
	}
	jsonData, _ := json.Marshal(datas)
	return jsonData
}

func HSelectTcpnet(hostName string,startTime,endTime int64) []byte {
	db := utils.SqlxCli()
	defer db.Close()
	var data []dataTypeStruck.TcpNetRespone

	sql := fmt.Sprintf("select hostName,allConn,timeUnix from monkey_s_tcpnetdata where hostName = '%v' and timeUnix >= %v and timeUnix < %v",hostName,startTime,endTime)
	if err := db.Select(&data,sql);err != nil{
		logUtils.Errorf("tvSelectTcpnet keyName=%v,startTime=%v,error=%v",hostName,startTime,err)
		return nil
	}
	jsonData, _ := json.Marshal(data)
	return jsonData
}
