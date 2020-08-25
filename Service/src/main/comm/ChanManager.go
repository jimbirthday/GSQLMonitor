package comm

import (
	"Gacos/src/main/base/domain"
	"errors"
	"log"
	"strconv"
	"time"
)

const (
	PREFIX        = "MSQL"
	DEFAULT_GROUP = "DEFAULT_GROUP"
)


func ListenDataStore() {
	for {
		select {
		case ack := <-DataChan:
			go putOrRepalce(ack)
		default:
			time.Sleep(time.Second * 10)
			log.Println("listenDataStore beat")
		}

	}

}

func putOrRepalce(service *domain.Service) {
	if _, ok := DataStore[service.NamespaceId+service.Name]; ok {
		delete(DataStore, service.NamespaceId+service.Name)
	}

	DataStore[service.NamespaceId+service.Name] = service
	log.Println("Client create serviceName:" + service.Name)
}


func CreateService(namespaceId string, servcieName string, formMap map[string]string) (service *domain.Service, err error) {
	s := new(domain.Service)

	s.Token = CreateToken(namespaceId, servcieName)
	s.CheckNums = 0
	s.NamespaceId = namespaceId
	s.Name = servcieName

	var defaultName = DEFAULT_GROUP
	if groupName, ok := formMap["groupName"]; ok {
		defaultName = groupName
	}

	var ip = ""
	if vIp, ok := formMap["ip"]; ok {
		ip = vIp
	} else {
		return nil, errors.New("ip can not empty")
	}

	var atoi = 0
	if port, ok := formMap["port"]; ok {

		atoi, err = strconv.Atoi(port)

		if err != nil {
			return nil, errors.New("port must be int type ")
		}

	} else {
		return nil, errors.New("port can not empty")
	}

	instance := new(domain.Instance)
	instance.Ip = ip
	instance.Port = atoi
	instance.ServiceName = servcieName
	instance.ClusterName = defaultName

	cluster := new(domain.Cluster)
	cluster.ClustetMap = make(map[string]*domain.Instance)
	cluster.ClustetMap[servcieName] = instance

	s.ClustetMap = make(map[string]*domain.Cluster)
	s.ClustetMap[defaultName] = cluster

	m := make(map[string]*domain.Service)
	m[servcieName] = service

	if ServiceMap == nil {
		ServiceMap = make(map[string]map[string]*domain.Service)
	}

	ServiceMap[namespaceId] = m

	return s, nil
}

func CreateToken(namespaceId string, servcieName string) string {
	return PREFIX + namespaceId + servcieName
}
