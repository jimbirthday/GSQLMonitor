package comm

import (
	"Gacos/src/main/base/domain"
	"log"
)

const (
	PREFIX        = "MSQL"
	DEFAULT_GROUP = "DEFAULT_GROUP"
)

func Init() {
	AddChan = make(chan *domain.Service, 10)
	RemoveChan = make(chan *domain.Service, 10)
	go listenAdd()
	go listenRemove()

}

func listenAdd() {

	for {

		select {

		case service := <-AddChan:
			go putOrRepalce(service)
		}

	}

}

func listenRemove() {

	for {

		select {

		case service := <-RemoveChan:
			go remove(service)

		}

	}

}

func remove(service *domain.Service) {

	if _, ok := DataStore[service.NamespaceId+service.Name]; ok {
		delete(DataStore, service.NamespaceId+service.Name)
	}

	log.Println("Client delete serviceName:" + service.Name)
}

func putOrRepalce(service *domain.Service) {

	if _, ok := DataStore[service.NamespaceId+service.Name]; ok {
		delete(DataStore, service.NamespaceId+service.Name)
	}

	DataStore[service.NamespaceId+service.Name] = service

	log.Println("Client create serviceName:" + service.Name)
}
