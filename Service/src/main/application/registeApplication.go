package application

import (
	baseWeb "Gacos/src/main/base/web"
	"Gacos/src/main/comm"
	"Gacos/src/main/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {
	comm.Gin = gin.Default()
	comm.Gin.Use(middleware.Cors())
	userInit()
	go comm.ListenDataStore()
}

func userInit() {
	baseRouter := comm.Gin.Group("/service")
	{
		baseRouter.POST("/create", baseWeb.Create)
	}
}
