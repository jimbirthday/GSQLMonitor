package web

import (
	"Gacos/src/main/base/domain"
	"Gacos/src/main/comm"
	"errors"
)

func GetService(namespaceId string, serviceName string) (service *domain.Service, err error) {
	if m, ok := comm.ServiceMap[namespaceId]; ok {
		if m != nil {

			if m, ok := m[serviceName]; ok {

				if m != nil {
					return m, nil
				}

			}

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
	for _, v := range comm.ServiceMap {

		if m, ok := v[service.Name]; ok {

			if m != nil {
				delete(v, service.Name)
			}

		}

	}

	go func() {

		comm.RemoveChan <- service

	}()
}
