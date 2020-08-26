package application

import (
	baseWeb "Gacos/src/main/base/web"
	"Gacos/src/main/comm"
	"Gacos/src/main/listener"
	"Gacos/src/main/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {
	comm.Gin = gin.Default()
	comm.Gin.Use(middleware.Cors())
	userInit()
	go comm.Init()
}

func userInit() {
	baseRouter := comm.Gin.Group("/service")
	{
		baseRouter.POST("/create", baseWeb.Create)
		baseRouter.POST("/remove", baseWeb.Remove)
		baseRouter.GET("/detail", baseWeb.Detail)
		baseRouter.GET("/list", baseWeb.List)
		baseRouter.GET("/heartbeat", listener.Heartbeat)
	}
}
