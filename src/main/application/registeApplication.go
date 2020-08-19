package application

import (
	"github.com/gin-gonic/gin"
	"middleProject/src/main/base/service"
	baseWeb "middleProject/src/main/base/web"
	"middleProject/src/main/comm"
	"middleProject/src/main/middleware"
	"net/http"
)

func Init() {
	comm.Gin = gin.Default()
	comm.Gin.Use(middleware.Cors())
	userInit()
}

func userInit() {
	baseRouter := comm.Gin.Group("/base")
	{
		baseRouter.GET("/processlist", service.Processlist)
		baseRouter.GET("/qps", service.QPS)
		baseRouter.GET("/tps", service.TPS)
		baseRouter.GET("/keyBuffer", service.BufferCache)
		baseRouter.GET("/innodbBuffer", service.InnoDBBufferCache)
		baseRouter.GET("/queryCache", service.QueryCache)
		baseRouter.GET("/threadCache", service.ThreadCache)
		baseRouter.GET("/lockStatus", service.LockStatus)
		baseRouter.GET("/dataBaseStatus", service.DataBaseStatus)
	}

	actionRouter := comm.Gin.Group("/action")
	{
		actionRouter.POST("/slowSql", service.SlowSQl)

	}

	comm.Gin.GET("/", baseWeb.IndexHandler)
	comm.Gin.LoadHTMLFiles("src/main/views/index.html")
	comm.Gin.StaticFS("/css", http.Dir("src/main/views/css"))
	comm.Gin.StaticFS("/js", http.Dir("src/main/views/js"))
	comm.Gin.StaticFS("/img", http.Dir("src/main/views/img"))

}
