package web

import (
	"github.com/gin-gonic/gin"
	"middleProject/src/main/base/service"
	"net/http"
)

func DataBaseStatus(c *gin.Context) {
	dbInfo := service.DataBaseStatus()
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": dbInfo,
		},
	)
}

//查看MySQl连接数
func Processlist(c *gin.Context) {
	processlist, err := service.Processlist()
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": http.StatusInternalServerError,
				"data": processlist,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": processlist,
		},
	)
}

//查询数据库QPS【每秒处理的查询数】
func QPS(c *gin.Context) {
	qps, err := service.QPS()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"data": err,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": qps,
		},
	)
}

//查看数据库TPS 【每秒处理的事务数】
func TPS(c *gin.Context) {
	tps, err := service.TPS()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"data": err,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": tps,
		},
	)

}

//key_buffer_read_hits = (1-key_reads / key_read_requests) * 100%
//key_buffer_write_hits = (1-key_writes / key_write_requests) * 100%
//key Buffer 命中率
func BufferCache(c *gin.Context) {
	buffer, err := service.BufferCache()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"data": err,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": buffer,
		},
	)
}

//InnoDB Buffer命中率
//innodb_buffer_read_hits = (1 - innodb_buffer_pool_reads / innodb_buffer_pool_read_requests) * 100%
func InnoDBBufferCache(c *gin.Context) {
	innodb, err := service.InnoDBBufferCache()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"data": err,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": innodb,
		},
	)
}

//Query Cache命中率
//Query_cache_hits = (Qcahce_hits / (Qcache_hits + Qcache_inserts )) * 100%;
func QueryCache(c *gin.Context) {
	query, err := service.QueryCache()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"data": err,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": query,
		},
	)
}

//Thread Cache 命中率
//Thread_cache_hits = (1 - Threads_created / connections ) * 100%
func ThreadCache(c *gin.Context) {
	thread, err := service.ThreadCache()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"data": err,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": thread,
		},
	)
}

//锁定状态
func LockStatus(c *gin.Context) {
	lock, err := service.LockStatus()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"data": err,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": lock,
		},
	)

}

func SlowSQl(c *gin.Context) {
	sql := c.PostForm("sql")
	res, err := service.SlowSQl(sql)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"data": err,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": res,
		},
	)

}
