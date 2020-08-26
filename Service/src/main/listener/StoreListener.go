package listener

import (
	"Gacos/src/main/comm"
	"fmt"
	"time"
)

func init() {
	go storeListener()
}

func storeListener() {

	for {
		time.Sleep(3 * time.Second)
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
					fmt.Print("remove ~~~~~~~~~~~~~~~~~~~~")
					comm.RemoveService(v)
				}

			}

		}

	}
}
