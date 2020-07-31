package dashboard

import (
	"encoding/json"
	"monkeyServer/dao"
	"monkeyServer/dataTypeStruck"
	"monkeyServer/utils"
)

type AreaInfo struct {
	City string `json:"city"`
}

func (c *AreaInfo) Get() dataTypeStruck.AreaRespone {


	var Respone dataTypeStruck.AreaRespone
	Respone.Data = make(map[string]interface{})


	startTime,endTime := utils.TvHourTimeUnix()
	sData := dao.AreaGet(c.City,startTime,endTime)

	for k, v := range sData {

		if k == "cpu" || k == "mem" || k == "tcp" {
			var ResponeData []dataTypeStruck.Respone
			_ = json.Unmarshal(v,&ResponeData)

			var data  = make(map[string][]dataTypeStruck.Respone)

			for _, v := range ResponeData {
				if _, ok := data[v.HostName];ok{
					data[v.HostName] = append(data[v.HostName],v)
				}else {
					data[v.HostName] = []dataTypeStruck.Respone{v}
				}
			}
			Respone.Data[k] = data
		}else {
			var diskData []dataTypeStruck.DiskRespone
			_ = json.Unmarshal(v,&diskData)
			DataRespone := make(map[string][]dataTypeStruck.DiskRespone)
			for _,v := range diskData {
				if _, ok := DataRespone[v.HostName]; ok{
					DataRespone[v.HostName] = append(DataRespone[v.HostName],v)
				}else {
					DataRespone[v.HostName] = []dataTypeStruck.DiskRespone{v}
				}
			}
			Respone.Data[k] = DataRespone
		}
	}
	return Respone
}

func AreaGetInfo(data []byte) dataTypeStruck.AreaRespone {
	var resquestData AreaInfo

	_ = json.Unmarshal(data,&resquestData)
	return resquestData.Get()
}