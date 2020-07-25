package dao

import (
	"encoding/json"
	"fmt"
	"monkeyServer/dataTypeStruck"
	"monkeyServer/logUtils"
	"monkeyServer/utils"
)




type Request = dataTypeStruck.TvRequest
//type Respone = dataTypeStruck.TvCpuRespone

func Get(c Request) []byte {
	switch c.Type {
	case 1:

		return tvSelectCpu(c.KeyName,c.StartTime,c.EndTime)
	case 2:
		return tvSelectMem(c.KeyName,c.StartTime,c.EndTime)
	case 3:
		return tvSelectTcpnet(c.KeyName,c.StartTime,c.EndTime)
	case 4:
		return tvSelectDisk(c.KeyName,c.StartTime,c.EndTime)
	}
	return nil
}


func tvSelectCpu(keyName string,startTime,endTime int64) []byte {


	db := utils.SqlxCli()
	defer db.Close()
	var data []dataTypeStruck.CpuRespone

	sql := fmt.Sprintf("select hostName,usedCpu,timeUnix from monkey_s_cpudata where keyName = '%v' and timeUnix >= %v and timeUnix < %v",keyName,startTime,endTime)
	logUtils.Debugf("tvSelectCpu sql=%v",sql)

	if err := db.Select(&data,sql);err != nil{
		logUtils.Errorf("TvSelectCpu keyName=%v,startTime=%v,error=%v",keyName,startTime,err)
		return nil
	}
	jsonData, _ := json.Marshal(data)
	return jsonData
}

func tvSelectMem(keyName string,startTime,endTime int64) []byte {
	db := utils.SqlxCli()
	defer db.Close()
	var data []dataTypeStruck.MemRespone

	sql := fmt.Sprintf("select hostName,used,timeUnix from monkey_s_memdata where keyName = '%v' and timeUnix >= %v and timeUnix < %v",keyName,startTime,endTime)
	if err := db.Select(&data,sql);err != nil{
		logUtils.Errorf("tvSelectMem keyName=%v,startTime=%v,error=%v",keyName,startTime,err)
		return nil
	}
	jsonData, _ := json.Marshal(data)
	return jsonData


}

func tvSelectDisk(keyName string,startTime,endTime int64) []byte {
	db := utils.SqlxCli()
	defer db.Close()
	var dbData []dataTypeStruck.DiskDB
	var datas []dataTypeStruck.DiskRespone

	sql := fmt.Sprintf("select hostName,disk,timeUnix from monkey_s_diskdata where keyName = '%v' and timeUnix >= %v and timeUnix < %v",keyName,startTime,endTime)
	if err := db.Select(&dbData,sql);err != nil{
		logUtils.Errorf("tvSelectDisk keyName=%v,startTime=%v,error=%v",keyName,startTime,err)
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

func tvSelectTcpnet(keyName string,startTime,endTime int64) []byte {
	db := utils.SqlxCli()
	defer db.Close()
	var data []dataTypeStruck.TcpNetRespone

	sql := fmt.Sprintf("select hostName,allConn,timeUnix from monkey_s_tcpnetdata where keyName = '%v' and timeUnix >= %v and timeUnix < %v",keyName,startTime,endTime)
	if err := db.Select(&data,sql);err != nil{
		logUtils.Errorf("tvSelectTcpnet keyName=%v,startTime=%v,error=%v",keyName,startTime,err)
		return nil
	}
	jsonData, _ := json.Marshal(data)
	return jsonData
}

func TvSelectStendData(tableName, hostName string,timeUnix int64) dataTypeStruck.StandData {
	db := utils.SqlxCli()
	defer db.Close()
	var data dataTypeStruck.StandData
	sql := fmt.Sprintf("select hostName,keyName,maValue,miValue,timeUnix from %v where hostName = '%v' and timeUnix = %v",tableName,hostName,timeUnix)
	logUtils.Debugf("TvSelectStendData sql=",sql)
	fmt.Println(sql)

	if err := db.Get(&data, sql); err != nil {
		logUtils.Error("tvSelectStendData tableName=%v,hostName=%v,timeUnix=%v,error=%v", tableName,hostName,timeUnix, err)
	}

	return data
}
