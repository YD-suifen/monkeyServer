package Calculation

import (
	"math"
	"monkeyServer/dao"
	"monkeyServer/dataTypeStruck"
	"monkeyServer/logUtils"
	"monkeyServer/utils"
	"monkeyServer/dao/calculation"

)

type StandData dataTypeStruck.StandData
var cpuTable string
var memTable string
var netTable string


func Algorithm()  {

	cpuTable = calculation.CreateTable("standCpu")
	memTable = calculation.CreateTable("standMem")
	netTable = calculation.CreateTable("standNet")
	timeList := utils.SeveDayUnix()
	for _, v := range timeList {
		for _,k := range v{
			algorithmCPU(k)
			algorithmMEM(k)
			algorithmNET(k)
		}
	}
}


func algorithmCPU(timeUnix int64)  {

	var source_cpu_list []float64
	var Stand StandData
	data := calculation.CpuCa(timeUnix)
	Sum := 0.0
	for _, v := range data{
		source_cpu_list = append(source_cpu_list,v.UsedCpu)
		Sum += v.UsedCpu
	}
	Stand.Update(data[0].HostName,data[0].KeyName,source_cpu_list,Sum,timeUnix)
	Stand.Insert(cpuTable)

}
func algorithmMEM(timeUnix int64)  {

	var source_mem_list []float64
	var Stand StandData
	data := calculation.MemCa(timeUnix)
	Sum := 0.0
	for _, v := range data{
		source_mem_list = append(source_mem_list,v.Used)
		Sum += v.Used
	}
	Stand.Update(data[0].HostName,data[0].KeyName,source_mem_list,Sum,timeUnix)
	Stand.Insert(memTable)

}
func algorithmNET(timeUnix int64)  {

	var source_cpu_list []float64
	var Stand StandData
	data := calculation.TcpnetCa(timeUnix)
	Sum := 0.0
	for _, v := range data{
		source_cpu_list = append(source_cpu_list,v.AllConn)
		Sum += v.AllConn
	}
	Stand.Update(data[0].HostName,data[0].KeyName,source_cpu_list,Sum,timeUnix)
	Stand.Insert(netTable)

}

func (c *StandData) Update(hostName,keyName string,data []float64,sum float64,timeUnix int64)  {

	source_cpu_len := float64(len(data))
	source_cpu_avg := sum / source_cpu_len

	var variance float64
	for _, v := range data {
		variance += math.Pow(v - source_cpu_avg,float64(2))
	}
	q := math.Sqrt(variance / source_cpu_len)
	cpuMax := source_cpu_avg + q * float64(3)
	cpuMin := source_cpu_avg - q * float64(3)
	c.KeyName = keyName
	c.HostName = hostName
	c.TimeUnix = timeUnix
	c.MaxValue = cpuMax
	c.MinValue = cpuMin
}

func (c *StandData) Insert(tableName string) {
	if calculation.Insert(tableName,c.KeyName,c.MaxValue,c.MinValue,c.TimeUnix){
		logUtils.Info("Insert true")
	}else {
		logUtils.Info("Insert false")
	}

}
