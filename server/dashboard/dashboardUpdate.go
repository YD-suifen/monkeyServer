package dashboard

import (
	"encoding/json"
	"monkeyServer/dao"
	"monkeyServer/dataTypeStruck"
	"monkeyServer/judgment/trend"
	"monkeyServer/utils"
)


type RequestTred dataTypeStruck.TrendDashboardRequest
type Respone dataTypeStruck.Respone

func Trend(jsonData []byte) dataTypeStruck.TrendDashboardRespone {

	var requestData RequestTred
	_ =json.Unmarshal(jsonData,&requestData)
	return requestData.TrendRequest()
}


func (c *RequestTred) TrendRequest() dataTypeStruck.TrendDashboardRespone {

	var ResponeData []dataTypeStruck.Respone
	var data dataTypeStruck.TrendDashboardRespone
	var diskData []dataTypeStruck.DiskRespone
	var DataRespone map[string][]dataTypeStruck.TrendRespone

	StartTime,EndTime := utils.TvHourTimeUnix()
	dbR := dataTypeStruck.TvRequest{
		c.City,
		c.Id,
		StartTime,
		EndTime,
	}
	if c.Id == 3 {
		ResponeByte := dao.Get(dbR)
		_ = json.Unmarshal(ResponeByte,&ResponeData)
		data.Data = ResponeData
		return data

	} else if c.Id == 4 {
		ResponeByte := dao.Get(dbR)
		_ = json.Unmarshal(ResponeByte,&diskData)
		data.Data = diskData
		return data
	}
	ResponeByte := dao.Get(dbR)
	_ = json.Unmarshal(ResponeByte,&ResponeData)
	dataList := trend.TrendActive(ResponeData,c.Id)
	DataRespone = make(map[string][]dataTypeStruck.TrendRespone)
	for _,v := range dataList {
		if _, ok := DataRespone[v.HostName]; ok{
			DataRespone[v.HostName] = append(DataRespone[v.HostName],v)
		}else {
			DataRespone[v.HostName] = []dataTypeStruck.TrendRespone{v}
		}

	}

	data.Data = DataRespone

	return data
}


//type dbRequest dataTypeStruck.TvRequest

//
//func (c *RequestTred) TrendRequest() dataTypeStruck.TrendDashboardRespone {
//	StartTime,EndTime := utils.TvHourTimeUnix()
//	switch c.Id {
//	case 1:
//		var dbR dataTypeStruck.TvRequest
//		var jsonData dataTypeStruck.TrendDashboardRespone
//		dbR.KeyName = c.City
//		dbR.Type = 1
//		dbR.StartTime = StartTime
//		dbR.EndTime = EndTime
//		list := dao.Get(dbR)
//		data := trend.Cpu(list)
//		jsonData.Data = data
//		return jsonData
//
//	case 2:
//	case 3:
//	case 4:
//
//
//	}
//
//
//}