package application

import (
	"github.com/gin-gonic/gin"
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
		baseRouter.GET("/processlist", baseWeb.Processlist)
		baseRouter.GET("/qps", baseWeb.QPS)
		baseRouter.GET("/tps", baseWeb.TPS)
		baseRouter.GET("/keyBuffer", baseWeb.BufferCache)
		baseRouter.GET("/innodbBuffer", baseWeb.InnoDBBufferCache)
		baseRouter.GET("/queryCache", baseWeb.QueryCache)
		baseRouter.GET("/threadCache", baseWeb.ThreadCache)
		baseRouter.GET("/lockStatus", baseWeb.LockStatus)
		baseRouter.GET("/dataBaseStatus", baseWeb.DataBaseStatus)
	}

	actionRouter := comm.Gin.Group("/action")
	{
		actionRouter.POST("/explainSql", baseWeb.ExplainSql)
		actionRouter.POST("/slowStatus", baseWeb.SlowStatus)
		actionRouter.POST("/slowInfo", baseWeb.SlowInfo)

	}

	comm.Gin.GET("/", baseWeb.IndexHandler)
	comm.Gin.LoadHTMLFiles("src/main/views/index.html")
	comm.Gin.StaticFS("/css", http.Dir("src/main/views/css"))
	comm.Gin.StaticFS("/js", http.Dir("src/main/views/js"))
	comm.Gin.StaticFS("/img", http.Dir("src/main/views/img"))

}
