package dataTypeStruck

type TvDataStruct struct {
	Return []HostData `json:"return"`
}
type HostData struct {
	HostName string `json:"hostName"`
	TypeName string `json:"typeName"`
}



type TvDashBoardTrend struct {
	TypeValue int `json:"typeValue"`
	Data interface{} `json:"data"`
}

type DataStruck struct {
	HostName string `json:"hostName"`
	KeyName string `json:"keyName"`
	Used float64 `json:"used"`
	State bool `json:"state"`
}

type DataDiskStruck struct {
	Disks []Disk `json:"disks"`
}




//type TrendDashBoardRespone struct {
//	TypeValue int `json:"typeValue"`
//	Data interface{} `json:"data"`
//}