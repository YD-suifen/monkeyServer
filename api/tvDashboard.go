package api

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"monkeyServer/server/dashboard"
	"monkeyServer/server/alarm"
	"net/http"
)



func Trend(c *gin.Context)  {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{"error":err})
	}else {
		jsonData := dashboard.Trend(data)
		c.JSON(http.StatusOK, jsonData)
	}
}

func Alarm(c *gin.Context) {

	if ok, data := alarm.ReadAlarmInfo();ok{
		c.JSON(http.StatusOK,data)
	}else {
		c.JSON(http.StatusOK,gin.H{"state":false})
	}
}
