package dashboard

import (
	"encoding/json"
	"monkeyServer/dao"
	"monkeyServer/dataTypeStruck"
	"monkeyServer/utils"
)

type HostInfo struct {
	HostName string `json:"hostName"`
}

func HostGetInfo(data []byte)  HostInfoS {

	var datas HostInfo

	_ = json.Unmarshal(data,&datas)

	return datas.Get()

}

type HostInfoS struct {
	Cpu interface{}
	Mem interface{}
	Disk interface{}
	Tcp interface{}
}

func (c *HostInfo) Get() HostInfoS {

	var Respone HostInfoS
	startTime,endTime := utils.TvHourTimeUnix()
	sData := dao.HostAll(c.HostName,startTime,endTime)

	for k, v := range sData {

		if k == "cpu" || k == "mem" || k == "tcp" {
			var ResponeData []dataTypeStruck.Respone
			_ = json.Unmarshal(v,&ResponeData)

			switch k {
			case "cpu":
				Respone.Cpu = ResponeData
				break
			case "mem":
				Respone.Mem = ResponeData
				break
			case "tcp":
				Respone.Tcp = ResponeData
			}

		}else {
			var diskData []dataTypeStruck.DiskRespone
			_ = json.Unmarshal(v,&diskData)

			Respone.Disk = diskData
		}
	}

	return Respone
}