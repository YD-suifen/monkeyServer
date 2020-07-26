package router

import (
	"github.com/gin-gonic/gin"
	"monkeyServer/api"
)



func RegistRouter()  {

	r := gin.Default()

	dashboard := r.Group("/dashboard")
	{
		dashboard.POST("/trend",api.Trend)
		dashboard.GET("/alarm",api.Alarm)
	}
	r.Run(":9534")
}