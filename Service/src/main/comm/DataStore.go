package comm

import "Gacos/src/main/base/domain"

var (
	DataStore map[string]*domain.Service
	AddChan  chan *domain.Service
	RemoveChan  chan *domain.Service
)

func init(){
	DataStore = make(map[string]*domain.Service)
}
