package dashboard

import (
	"encoding/json"
	"fmt"
	"monkeyServer/dao"
	"monkeyServer/dataTypeStruck"
	"monkeyServer/utils"
)



type DashboardPost struct {
	City string `json:"city"`
	Id int `json:"Id"`
}



func (c *DashboardPost) Update() []byte {

	var request dataTypeStruck.TvRequest
	startTime,endTime := utils.TvHourTimeUnix()
	request.KeyName = c.City
	request.Type = c.Id
	request.StartTime = startTime
	request.EndTime = endTime
	return dao.Get(request)

}


func Dashboard(jsonData []byte) []byte {
	fmt.Println("cc",string(jsonData))
	var requestData DashboardPost
	err := json.Unmarshal(jsonData,&requestData)
	if err != nil{
		fmt.Println("errr=",err)
	}
	fmt.Println("ddd:",requestData.Id)
	return requestData.Update()
}
