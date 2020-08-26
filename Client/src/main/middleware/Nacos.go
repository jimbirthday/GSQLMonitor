package middleware

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Service struct {
	NamespaceId string
	ServiceName string
	GroupName   string
	Ip          string
	Port        string
}

var (
	service       Service
	ServiceCreate bool
)

func init() {
	go CreateService()
}
func CreateService() {
	m := make(map[string][]string)
	namespaceId := []string{"SPACE"}
	serviceName := []string{"LOCAL_SERVICE"}
	ip := []string{getIp()}
	port := []string{"8008"}
	m["namespaceId"] = namespaceId
	m["serviceName"] = serviceName
	m["ip"] = ip
	m["port"] = port

	err := postFrom("http://localhost:8002/service/create", m)
	if err != nil {
		ServiceCreate = false
		go retryCreateService()
	}

}

func retryCreateService() {

	if ServiceCreate {
		return
	}

	time.Sleep(3 * time.Second)

	fmt.Println("retry create service ")

	go CreateService()
}

func postFrom(path string, value map[string][]string) error {
	resp, err := http.PostForm(path,
		value)
	if err != nil {
		// handle error
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return err
	}
	fmt.Println(string(body))
	return nil
}

func getIp() string {
	var ip = ""
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("get ip failed")
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				fmt.Println(ipnet.IP.String())
			}

		}
	}
	return ip
}
