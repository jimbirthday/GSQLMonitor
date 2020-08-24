package web

import "github.com/gin-gonic/gin"

func IndexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
