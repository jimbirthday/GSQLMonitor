package listener

import (
	"Gacos/src/main/comm"
	"time"
)

func init() {
	go storeListener()
}

func storeListener() {

	for {
		time.Sleep(30 * time.Second)
		go checkHealth()
	}

}

func checkHealth() {

	currtTime := time.Now().Unix()
	for _, v := range comm.ServiceMap {

		for _, v := range v {

			for _, v := range v {
				i := currtTime - v.CheckTime

				if i > 30 {
					comm.RemoveService(v)
				}

			}

		}

	}
}
