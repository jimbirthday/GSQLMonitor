package web

import (
	"Gacos/src/main/base/domain"
	"Gacos/src/main/comm"
	"errors"
)

func GetService(namespaceId string, serviceName string) (service *domain.Service, err error) {
	if m, ok := comm.ServiceMap[namespaceId]; ok {

		if m, ok := m[serviceName]; ok {
			return m, nil
		}

	}
	return nil, errors.New("the service is not exist ")
}

func addOrReplaceService(service *domain.Service) {
	go func() {
		comm.DataChan <- service
	}()
}
