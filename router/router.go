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
	}
	r.Run(":9534")
}