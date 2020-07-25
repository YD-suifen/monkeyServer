package dataTypeStruck


type HostInfo struct {
	HostName string `json:"hostName" db:"hostName"`
	TimeUnix int64 `json:"timeUnix" db:"timeUnix"`
}

type CpuRespone struct {
	HostInfo
	Used float64 `json:"used" db:"usedCpu"`
}

type MemRespone struct {
	HostInfo
	Used float64 `json:"used" db:"used"`
}

type TcpNetRespone struct {
	HostInfo
	Used int `json:"used" db:"allConn"`
}

type Respone struct {
	HostInfo
	Used float64 `json:"used"`
}


type DiskRespone struct {
	HostInfo
	Disks []Disk `json:"disks"`
}

type DiskDB struct {
	HostInfo
	Disk string `json:"disk" db:"disk"`
}

type Disk struct {
	DevName string `json:"devName" db:"devName"`
	Total float64 `json:"total" db:"total"`
	Used float64 `json:"used" db:"used"`
	Free float64 `json:"free" db:"free"`
}



type TvRequest struct {
	KeyName string `json:"keyName"`
	Type int `json:"type"`
	StartTime int64 `json:"startTime"`
	EndTime int64 `json:"endTime"`
}


type TrendDashboardRequest struct {
	City string `json:"city"`
	Id int `json:"Id"`
}



type TrendRespone struct {
	HostInfo
	Used float64 `json:"used"`
	State int `json:"state"`
}

type TrendDashboardRespone struct {
	Data interface{} `json:"data"`
}














//type CpuRespone struct {
//	HostName string `json:"hostName" db:"hostName"`
//	UsedCpu float64 `json:"usedCpu" db:"usedCpu"`
//	TimeUnix int64 `json:"timeUnix" db:"timeUnix"`
//}
//
//type MemRespone struct {
//	HostName string `json:"hostName" db:"hostName"`
//	Used float64 `json:"used" db:"used"`
//	TimeUnix int64 `json:"timeUnix" db:"timeUnix"`
//}
//
//type DiskRespone struct {
//	HostName string `json:"hostName"`
//	Disks []Disk `json:"disks"`
//	TimeUnix int64 `json:"timeUnix"`
//}
//
//type DiskDB struct {
//	HostName string `json:"hostName" db:"hostName"`
//	Disk string `json:"disk" db:"disk"`
//	TimeUnix int64 `json:"timeUnix" db:"timeUnix"`
//}
//
//type Disk struct {
//	DevName string `json:"devName" db:"devName"`
//	Total float64 `json:"total" db:"total"`
//	Used float64 `json:"used" db:"used"`
//	Free float64 `json:"free" db:"free"`
//}
//
//
//
//type TcpNetRespone struct {
//	HostName string `json:"hostName" db:"hostName"`
//	AllConn int `json:"allConn" db:"allConn"`
//	TimeUnix int64 `json:"timeUnix" db:"timeUnix"`
//}

