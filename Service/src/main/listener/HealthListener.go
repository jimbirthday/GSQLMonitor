package listener

import (
	"Gacos/src/main/base/domain"
	"Gacos/src/main/comm"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var (
	HeartBeatChanMap map[string]chan *domain.Service
	HeartBeatChan    chan string
)

func init() {
	HeartBeatChanMap = make(map[string]chan *domain.Service)
	HeartBeatChan = make(chan string, 10)
	go ListerHeart()
}

func Heartbeat(c *gin.Context) {

	namespaceId := c.Query("namespaceId")
	serviceName := c.Query("serviceName")

	go func() {
		HeartBeatChan <- namespaceId + "," + serviceName
	}()

	c.JSON(
		http.StatusOK,
		gin.H{
		},
	)
}

func ListerHeart() {
	for {

		select {

		case name := <-HeartBeatChan:
			split := strings.Split(name, ",")

			servicesChan := HeartBeatChanMap[split[0]+split[1]]
			if servicesChan == nil {

				servicesChan = make(chan *domain.Service, 10)
				HeartBeatChanMap[split[0]+split[1]] = servicesChan
			}

			go HeartbeatCheckIn(servicesChan)
			go HeartbeatCheckOut(servicesChan, split[0], split[1])
		}

	}
}

func HeartbeatCheckIn(ownChen chan *domain.Service) {
	ownChen <- nil
	return
}

func HeartbeatCheckOut(ownChen chan *domain.Service, namespaceId string, serviceName string) {

	<-ownChen

	service, _ := comm.GetService(namespaceId, serviceName)

	if service == nil {
		return
	}
	timeUnix := time.Now().Unix()
	service.CheckTime = timeUnix
	service.CheckNums = service.CheckNums + 1

	comm.AddOrReplaceService(service)

}
