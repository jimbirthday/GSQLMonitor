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

	service, err := GetService(namespaceId, serviceName)
	if service != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "specified service already exists, serviceName : " + serviceName,
			},
		)

		return
	}

	service, err = comm.CreateService(namespaceId, serviceName, formMap)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "service create failed , serviceName : " + serviceName,
			},
		)

		return
	}

	addOrReplaceService(service)

	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"msg":  "service create suc , serviceName : " + serviceName,
		},
	)
}

func Remove(c *gin.Context) {

	namespaceId := c.PostForm("namespaceId")
	serviceName := c.PostForm("serviceName")

	service, err := GetService(namespaceId, serviceName)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "specified service not exist, serviceName : " + serviceName,
			},
		)

		return
	}

	RemoveService(service)

	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"msg":  "service remove suc , serviceName : " + serviceName,
		},
	)

}

func Detail(c *gin.Context) {

	namespaceId := c.Query("namespaceId")
	serviceName := c.Query("serviceName")

	service, err := GetService(namespaceId, serviceName)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "specified service not exist, serviceName : " + serviceName,
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": service,
		},
	)

}

func List(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": comm.ServiceMap,
		},
	)

}
