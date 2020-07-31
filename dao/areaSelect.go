package dao

func AreaGet(keyName string,startTime,endTime int64) map[string][]byte {

	data := make(map[string][]byte)
	data["cpu"] = tvSelectCpu(keyName,startTime,endTime)
	data["mem"] = tvSelectMem(keyName,startTime,endTime)
	data["disk"] = tvSelectDisk(keyName,startTime,endTime)
	data["tcp"] = tvSelectTcpnet(keyName,startTime,endTime)
	return data
}

