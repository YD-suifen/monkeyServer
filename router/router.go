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
		dashboard.POST("/area",api.AreaDash)
		dashboard.POST("/host",api.HostDash)
	}
	r.Run(":9534")
}