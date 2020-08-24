package comm

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

var  (
	Gin *gin.Engine //gin.Default
	Db *xorm.Engine //数据库链接
	DbInfo *xorm.Engine //数据库链接
)
