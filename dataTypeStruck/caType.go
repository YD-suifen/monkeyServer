package dataTypeStruck

type Cacpu struct {
	HostName string `db:"hostName"`
	KeyName string `db:"keyName"`
	UsedCpu float64 `db:"usedCpu"`
	TimeUnix int64 `db:"timeUnix"`
}

type Camem struct {
	HostName string `db:"hostName"`
	KeyName string `db:"keyName"`
	Used float64 `db:"used"`
	TimeUnix int64 `db:"timeUnix"`
}


type Catcpnet struct {
	HostName string `db:"hostName"`
	KeyName string `db:"keyName"`
	AllConn float64 `db:"allConn"`
	TimeUnix int64 `db:"timeUnix"`
}


//type StandData struct {
//	HostName string `json:"hostName" db:"hostName"`
//	KeyName string `json:"keyName" db:"keyName"`
//	MaxValue float64 `json:"maxValue" db:"maValue"`
//	MinValue float64 `json:"minValue" db:"miValue"`
//	TimeUnix int64 `json:"timeUnix" db:"timeUnix"`
//}

type StandData struct {
	HostName string `db:"hostName"`
	KeyName string `db:"keyName"`
	MaxValue float64 `db:"maValue"`
	MinValue float64 `db:"miValue"`
	TimeUnix int64 `db:"timeUnix"`
}