package comm

import "Gacos/src/main/base/domain"

var (
	ServiceMap map[string]map[string]map[string]*domain.Service
)

func init() {
	ServiceMap = make(map[string]map[string]map[string]*domain.Service)
}
