package comm

import "Gacos/src/main/base/domain"

var (
	DataStore map[string]*domain.Service
	DataChan  chan *domain.Service
)
