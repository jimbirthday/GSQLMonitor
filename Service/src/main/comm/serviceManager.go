package comm

import (
	"Gacos/src/main/base/domain"
	"errors"
	"strconv"
)

func GetService(namespaceId string, serviceName string) (service *domain.Service, err error) {
	if m, ok := DataStore[namespaceId+serviceName]; ok {

		if m != nil {
			return m, nil
		}

	}
	return nil, errors.New("the service is not exist ")
}

func AddOrReplaceService(service *domain.Service) {
	go func() {

		AddChan <- service

	}()
}

func RemoveService(service *domain.Service) {
	if v := ServiceMap[service.NamespaceId]; v != nil {

		for _, v := range v {

			if v1 := v[service.Name]; v1 != nil {
				delete(v, service.Name)
			}

		}

	}

	go func() {

		RemoveChan <- service

	}()
}



func CreateService(namespaceId string, servcieName string, formMap map[string]string) (service *domain.Service, err error) {
	s := new(domain.Service)

	s.Token = CreateToken(namespaceId, servcieName)
	s.CheckNums = 0
	s.NamespaceId = namespaceId
	s.Name = servcieName

	var defaultName = DEFAULT_GROUP
	if groupName, ok := formMap["groupName"]; ok {

		if groupName != "" {
			defaultName = groupName
		}

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

	s.Ins = instance

	nameMap := ServiceMap[namespaceId]
	if nameMap == nil {
		culterMap := make(map[string]map[string]*domain.Service)

		serMap := make(map[string]*domain.Service)
		serMap[servcieName] = s
		culterMap[defaultName] = serMap

		ServiceMap[namespaceId] = culterMap

	} else {
		serMap := nameMap[defaultName]

		if serMap == nil {

			serMap = make(map[string]*domain.Service)
			serMap[defaultName] = s

			nameMap[defaultName] = serMap
		} else {
			serMap[servcieName] = s
		}
		ServiceMap[namespaceId] = nameMap
	}

	return s, nil
}

func CreateToken(namespaceId string, servcieName string) string {
	return PREFIX + namespaceId + servcieName
}
