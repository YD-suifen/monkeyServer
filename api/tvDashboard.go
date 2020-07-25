package api

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"monkeyServer/server/dashboard"
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
