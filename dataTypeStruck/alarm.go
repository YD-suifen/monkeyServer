package dataTypeStruck


type NoAlarmInfo struct {
	Data interface{}
	Type int
}

type NoAlarmCpu struct {
	CpuRespone
	KeyName string `json:"keyName" db:"keyName"`
}
type NoAlarmMem struct {
	MemRespone
	KeyName string `json:"keyName" db:"keyName"`
}
type NoAlarmTcp struct {
	TcpNetRespone
	KeyName string `json:"keyName" db:"keyName"`
}


type AlarmInfo struct {
	HostInfo
	KeyName string `json:"keyName"`
	Used float64
	Type int
	State bool
}