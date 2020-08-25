package web

import (
	"Gacos/src/main/base/domain"
	"Gacos/src/main/comm"
	"errors"
)

func GetService(namespaceId string, serviceName string) (service *domain.Service, err error) {
	if m, ok := comm.DataStore[namespaceId+serviceName]; ok {

		if m != nil {
			return m, nil
		}

	}
	return nil, errors.New("the service is not exist ")
}

func addOrReplaceService(service *domain.Service) {
	go func() {

		comm.AddChan <- service

	}()
}

func RemoveService(service *domain.Service) {
	if v := comm.ServiceMap[service.NamespaceId]; v != nil {

		for _, v := range v {

			if v1 := v[service.Name]; v1 != nil {
				delete(v, service.Name)
			}

		}

	}

	go func() {

		comm.RemoveChan <- service

	}()
}
