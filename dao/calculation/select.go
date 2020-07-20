package calculation

import (
	"fmt"
	"monkeyServer/dataTypeStruck"
	"monkeyServer/logUtils"
	"monkeyServer/utils"
)

func CpuCa(time int64) []dataTypeStruck.Cacpu {

	db := utils.SqlxCli()
	defer db.Close()

	var data []dataTypeStruck.Cacpu

	sql := fmt.Sprintf("select hostName,keyName,usedCpu,timeUnix from monkey_s_cpudata where timeUnix = %v",time)
	if err := db.Select(&data,sql);err != nil{
		logUtils.Errorf("CpuCa Time=%v,error=%v",time,err)
		return nil
	}
	return data
}


func MemCa(time int64) []dataTypeStruck.Camem {

	db := utils.SqlxCli()
	defer db.Close()

	var data []dataTypeStruck.Camem

	sql := fmt.Sprintf("select hostName,keyName,used,timeUnix from monkey_s_memdata where timeUnix = %v",time)
	if err := db.Select(&data,sql);err != nil{
		logUtils.Errorf("CpuCa Time=%v,error=%v",time,err)
		return nil
	}
	return data
}


func TcpnetCa(time int64) []dataTypeStruck.Catcpnet {

	db := utils.SqlxCli()
	defer db.Close()

	var data []dataTypeStruck.Catcpnet

	sql := fmt.Sprintf("select hostName,keyName,allConn,timeUnix from monkey_s_tcpnetdata where timeUnix = %v",time)
	if err := db.Select(&data,sql);err != nil{
		logUtils.Errorf("CpuCa Time=%v,error=%v",time,err)
		return nil
	}
	return data
}