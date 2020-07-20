package router

import (
	"github.com/gin-gonic/gin"
	"monkeyServer/api"
)


func RegistRouter(r *gin.Engine)  {

	r.POST("/dashboard", api.Dashboard)

}