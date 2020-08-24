package web

import (
	"Gacos/src/main/comm"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {

	namespaceId := c.PostForm("namespaceId")
	serviceName := c.PostForm("serviceName")

	ip := c.PostForm("ip")
	port := c.PostForm("port")
	groupName := c.PostForm("groupName")

	formMap := make(map[string]string)

	formMap["ip"] = ip
	formMap["port"] = port
	formMap["groupName"] = groupName

	_, err := GetService(namespaceId, serviceName)
	if err == nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"msg": "specified service already exists, serviceName : " + serviceName,
			},
		)

		return
	}

	service, err := comm.CreateService(namespaceId, serviceName, formMap)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"msg": "service create failed , serviceName : " + serviceName,
			},
		)

		return
	}

	addOrReplaceService(service)

	c.JSON(
		http.StatusOK,
		gin.H{
			"msg": "service create suc , serviceName : " + serviceName,
		},
	)
}
