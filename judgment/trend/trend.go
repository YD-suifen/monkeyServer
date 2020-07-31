package trend

import (
	"monkeyServer/dao"
	"monkeyServer/dataTypeStruck"
	"monkeyServer/utils"
	"strconv"
	"sync"
)

type State struct {
	Type string
	Count int
}



var DataMap map[string]State
var RWmutex *sync.RWMutex

func init()  {
	DataMap = make(map[string]State)
	RWmutex  = &sync.RWMutex{}
}


func TrendActive(sourceData []dataTypeStruck.Respone,Type int) []dataTypeStruck.TrendRespone {
	var tableName string
	var disRespone []dataTypeStruck.TrendRespone
	var disData dataTypeStruck.TrendRespone
	typeString := strconv.Itoa(Type)
	switch Type {
	case 1:
		tableName = "standCpu_" + utils.Yesterday()
		break
	case 2:
		tableName = "standMem_" + utils.Yesterday()
		break
	}

	//fmt.Println("sourceData len=",len(sourceData))

	for _, v := range sourceData{
		stendData := dao.TvSelectStendData(tableName,v.HostName,v.TimeUnix - 86400)

		disData.HostName = v.HostName
		disData.TimeUnix = v.TimeUnix
		disData.Used = v.Used

		RWmutex.Lock()
		defer RWmutex.Unlock()
		if i := av(v.Used,stendData.MinValue,stendData.MinValue); i == 0{

			delete(DataMap,v.HostName+"up"+typeString)
			delete(DataMap,v.HostName+"down"+typeString)
			disData.State = 0
			disRespone = append(disRespone,disData)

		}else if i == 1 {
			if k, ok := DataMap[v.HostName+"up"+typeString]; ok {
				k.Count += 1
			}else {
				DataMap[v.HostName+"up"+typeString] = State{"on",1}
			}
			cout := DataMap[v.HostName+"up"+typeString].Count
			disData.State = ac(cout)
			disRespone = append(disRespone,disData)
		}else if i == 2 {
			if k, ok := DataMap[v.HostName + "down" +typeString]; ok {
				k.Count += 1
			}else {
				DataMap[v.HostName + "down"+typeString] = State{"down",1}
			}
			cout := DataMap[v.HostName + "down"+typeString].Count
			disData.State = ad(cout)
			disRespone = append(disRespone,disData)
		}
	}
	//fmt.Println("disRespone len=",len(disRespone))
	return disRespone
}

func av(sour,stMax,stMin float64) int {
	i := 0
	if sour <= stMax && sour >= stMin {
		return i
	}else if sour > stMax {
		return i+1
	}else if sour < stMin{
		return i+2
	}
	return i
}

func ac(a int) int {
	if a > 2 && a <= 3 {
		return 1
	}else if a > 3 && a <= 5 {
		return 2
	}
	return 0
}
func ad(a int) int {
	if a > 2 && a <= 3 {
		return 3
	}else if a > 3 && a <= 5 {
		return 4
	}
	return 0
}

//func (c *Respone) Check(max,min float64) int {
//	return av(c.Used,max,min)
//}














//func Cpu(list []byte) []dataTypeStruck.TrendRespone {
//	var sourData []dataTypeStruck.TvCpuRespone
//	var DirData []dataTypeStruck.TrendRespone
//	_ = json.Unmarshal(list,&sourData)
//	for _, v := range sourData{
//		tableName := "standCpu_" + utils.Yesterday()
//		stendData := dao.TvSelectStendData(tableName,v.HostName,v.TimeUnix - 86400)
//		if i := av(v.UsedCpu,stendData.MaxValue,stendData.MinValue); i != 0{
//			if _, ok := CpuMap[v.HostName]; ok {
//				if i == 1 {
//					CpuMap[v.HostName].Count++
//				}else if i == 2{
//					CpuMap[v.HostName].Count++
//				}else if i == 0{
//					delete(CpuMap, v.HostName)
//				}
//			}else {
//				if i == 1 {
//					CpuMap[v.HostName] = State{"on",1}
//				}else if i == 2{
//					CpuMap[v.HostName] = State{"down",1}
//				}
//			}
//		}
//		data := dataTypeStruck.TrendRespone{}
//		if _, ok := CpuMap[v.HostName]; ok {
//			data.HostName = v.HostName
//			data.Used = v.UsedCpu
//			if CpuMap[v.HostName].Type == "on" && CpuMap[v.HostName].Count > 2 && CpuMap[v.HostName].Count < 4 {
//				data.State = 1
//			} else if CpuMap[v.HostName].Type == "on" && CpuMap[v.HostName].Count > 4 && CpuMap[v.HostName].Count < 6 {
//				data.State = 2
//			} else if CpuMap[v.HostName].Type == "on" && CpuMap[v.HostName].Count > 6 {
//				data.State = 3
//			} else if CpuMap[v.HostName].Type == "down" && CpuMap[v.HostName].Count > 2 && CpuMap[v.HostName].Count < 4 {
//				data.State = 4
//			} else if CpuMap[v.HostName].Type == "on" && CpuMap[v.HostName].Count > 4 && CpuMap[v.HostName].Count < 6 {
//				data.State = 5
//			} else if CpuMap[v.HostName].Type == "on" && CpuMap[v.HostName].Count > 6{
//				data.State = 6
//			}
//		} else {
//			data.HostName = v.HostName
//			data.Used = v.UsedCpu
//			data.State = 0
//		}
//		DirData = append(DirData,data)
//
//	}
//	return DirData
//}

