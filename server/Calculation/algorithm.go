package Calculation

import (
	"fmt"
	"math"
	"monkeyServer/dao/calculation"
	"monkeyServer/dataTypeStruck"
	"monkeyServer/logUtils"
	"monkeyServer/utils"
	"strconv"
	"time"
)

type StandData dataTypeStruck.StandData
var CpuTable string
var MemTable string
//var netTable string
var TodayTimeUnix int64


func Algorithm()  {

	CpuTable = calculation.CreateTable("standCpu")
	MemTable = calculation.CreateTable("standMem")
	//netTable = calculation.CreateTable("standNet")
	timeList := utils.SeveDayUnix()
	for _, v := range timeList {
		TodayTimeUnix = 0
		TodayTimeUnix = v[6] + 86400

		algorithmCPU1(TodayTimeUnix,v)
		algorithmMEM1(TodayTimeUnix,v)

		time.Sleep(time.Second * 1)
	}
}

type A struct {
	KeyName string `json:"keyName"`
	Count int `json:"count"`
	Avg float64 `json:"avg"`
	Used float64 `json:"used"`
	Variance float64 `json:"variance"`
	MaxValue float64 `json:"maxValue"`
	MinValue float64 `json:"minValue"`
	TimeUnix int64 `json:"timeUnix"`
}

func algorithmCPU1(TodayTimeUnix int64, arg []int64)  {

	var nTimeList []dataTypeStruck.Cacpu
	for _, v := range arg {
		data := calculation.CpuCa(v)
		nTimeList = append(nTimeList,data...)
	}
	//fmt.Println(len(nTimeList))
	HostInfo := make(map[string]*A)

	for _, v := range nTimeList {
		if s, ok := HostInfo[v.HostName]; ok{
			s.Used += v.UsedCpu
			s.KeyName = v.KeyName
			s.Count++
		}else {
			HostInfo[v.HostName] = &A{}
			HostInfo[v.HostName].KeyName = v.KeyName
			HostInfo[v.HostName].Used = v.UsedCpu
			HostInfo[v.HostName].Count = 1
		}
	}
	for k, v := range HostInfo {

		v.Avg = v.Used / float64(v.Count)
		for _, h := range nTimeList {
			if k == h.HostName {
				v.Variance += math.Pow(h.UsedCpu - v.Avg,float64(2))
			}
		}
		q := math.Sqrt(v.Variance / float64(v.Count))
		cpuMax := v.Avg + q * float64(3)
		cpuMin := v.Avg - q * float64(3)
		
		v.TimeUnix = TodayTimeUnix
		v.MaxValue = decimal(cpuMax)
		v.MinValue = decimal(cpuMin)
		if calculation.Insert(CpuTable,k,v.KeyName,v.MaxValue,v.MinValue,v.TimeUnix){
			logUtils.Info("Insert true")
		}else {
			logUtils.Info("Insert false")
		}
	}
}


func algorithmMEM1(TodayTimeUnix int64, arg []int64)  {

	var nTimeList []dataTypeStruck.Camem
	for _, v := range arg {
		data := calculation.MemCa(v)
		nTimeList = append(nTimeList,data...)
	}
	fmt.Println(len(nTimeList))
	HostInfo := make(map[string]*A)

	for _, v := range nTimeList {
		if s, ok := HostInfo[v.HostName]; ok{
			s.Used += v.Used
			s.KeyName = v.KeyName
			s.Count++
		}else {
			HostInfo[v.HostName] = &A{}
			HostInfo[v.HostName].KeyName = v.KeyName
			HostInfo[v.HostName].Used = v.Used
			HostInfo[v.HostName].Count = 1
		}
	}
	for k, v := range HostInfo {

		v.Avg = v.Used / float64(v.Count)
		for _, h := range nTimeList {
			if k == h.HostName {
				v.Variance += math.Pow(h.Used - v.Avg,float64(2))
			}
		}
		q := math.Sqrt(v.Variance / float64(v.Count))
		cpuMax := v.Avg + q * float64(3)
		cpuMin := v.Avg - q * float64(3)

		v.TimeUnix = TodayTimeUnix
		v.MaxValue = decimal(cpuMax)
		v.MinValue = decimal(cpuMin)
		if calculation.Insert(MemTable,k,v.KeyName,v.MaxValue,v.MinValue,v.TimeUnix){
			logUtils.Info("Insert true")
		}else {
			logUtils.Info("Insert false")
		}
	}
}
//func (c *StandData) Update(data A,timeUnix int64)  {
//
//	source_cpu_len := float64(len(data))
//	source_cpu_avg := sum / source_cpu_len
//
//	var variance float64
//	for _, v := range data {
//		variance += math.Pow(v - source_cpu_avg,float64(2))
//	}
//	q := math.Sqrt(variance / source_cpu_len)
//	cpuMax := source_cpu_avg + q * float64(2)
//	cpuMin := source_cpu_avg - q * float64(2)
//	c.KeyName = keyName
//	c.HostName = hostName
//	c.TimeUnix = timeUnix
//	c.MaxValue = decimal(cpuMax)
//	c.MinValue = decimal(cpuMin)
//}

//func algorithmCPU(timeUnix,todayUnix int64)  {
//
//	var source_cpu_list []float64
//	var Stand StandData
//	data := calculation.CpuCa(timeUnix)
//	if len(data) == 0{
//		return
//	}
//	Sum := 0.0
//	for _, v := range data{
//		source_cpu_list = append(source_cpu_list,v.UsedCpu)
//		Sum += v.UsedCpu
//	}
//	logUtils.Infof("algorithmCPU lendata=%v,lensource_cpu_list=%v,Sum=%v,cpuTable=%v",len(data),len(source_cpu_list),Sum,CpuTable)
//	Stand.Update(data[0].HostName,data[0].KeyName,source_cpu_list,Sum,todayUnix)
//	Stand.Insert(CpuTable)
//
//}
//func algorithmMEM(timeUnix,todayUnix int64)  {
//
//	var source_mem_list []float64
//	var Stand StandData
//	data := calculation.MemCa(timeUnix)
//	if len(data) == 0{
//		return
//	}
//	Sum := 0.0
//	for _, v := range data{
//		source_mem_list = append(source_mem_list,v.Used)
//		Sum += v.Used
//	}
//	Stand.Update(data[0].HostName,data[0].KeyName,source_mem_list,Sum,todayUnix)
//	Stand.Insert(MemTable)
//
//}




func (c *StandData) Insert(tableName string) {
	if calculation.Insert(tableName,c.HostName,c.KeyName,c.MaxValue,c.MinValue,c.TimeUnix){
		logUtils.Info("Insert true")
	}else {
		logUtils.Info("Insert false")
	}

}


func decimal(value float64) float64  {

	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

//monkey_s_cpudata
//monkey_s_diskdata
//monkey_s_memdata
//monkey_s_tcpnetdata
//
//ALTER TABLE `table_name` ADD INDEX index_name ( `column` )
//ALTER TABLE `table_name` ADD INDEX index_name ( `column1`, `column2`, `column3` )