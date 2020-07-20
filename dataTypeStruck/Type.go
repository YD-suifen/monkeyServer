package dataTypeStruck

type TvDataStruct struct {
	Return []HostData `json:"return"`
}
type HostData struct {
	HostName string `json:"hostName"`
	TypeName string `json:"typeName"`
}

type TvRequest struct {
	KeyName string `json:"keyName"`
	Type int `json:"type"`
	StartTime int64 `json:"startTime"`
	EndTime int64 `json:"endTime"`
}

type TvCpuRespone struct {
	HostName string `json:"hostName" db:"hostName"`
	UsedCpu float64 `json:"usedCpu" db:"usedCpu"`
	IdleCpu float64 `json:"idleCpu" db:"idleCpu"`
	TimeUnix int64 `json:"timeUnix" db:"timeUnix"`
}

type TvMemRespone struct {
	HostName string `json:"hostName" db:"hostName"`
	Total float64 `json:"total" db:"total"`
	Used float64 `json:"used" db:"used"`
	Free float64 `json:"free" db:"free"`
	TimeUnix int64 `json:"timeUnix" db:"timeUnix"`
}

type TvDiskDB struct {
	HostName string `json:"hostName" db:"hostName"`
	Disk string `json:"disk" db:"disk"`
	TimeUnix int64 `json:"timeUnix" db:"timeUnix"`
}

type Disk struct {
	DevName string `json:"devName" db:"devName"`
	Total float64 `json:"total" db:"total"`
	Used float64 `json:"used" db:"used"`
	Free float64 `json:"free" db:"free"`
}

type TvDiskRespone struct {
	HostName string `json:"hostName"`
	Disks []Disk `json:"disks"`
	TimeUnix int64 `json:"timeUnix"`
}

type TvTcpNetRespone struct {
	HostName string `json:"hostName" db:"hostName"`
	AllConn int `json:"allConn" db:"allConn"`
	Established int `json:"established" db:"established"`
	TimeUnix int64 `json:"timeUnix" db:"timeUnix"`
}