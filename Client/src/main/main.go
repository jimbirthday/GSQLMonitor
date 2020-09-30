package main

import (
	"middleProject/src/main/application"
	"middleProject/src/main/comm"
)

func main() {

}

func init() {
	application.Init()
	comm.InitDB()
	err := comm.Gin.Run("0.0.0.0:8009") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic("application run failed  err is " + err.Error())
	}
}
