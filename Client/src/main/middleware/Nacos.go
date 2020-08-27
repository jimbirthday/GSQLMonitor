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
	service       *Service
	ServiceCreate bool
)

const URL string = "http://localhost:8002"
const namespaceId string = "G1"
const serviceName string = "GSQL"
const port string = "8008"

func init() {
	go CreateService()
	go HealthCheck()
}
func CreateService() {

	m := make(map[string][]string)
	namespaceId := []string{namespaceId}
	serviceName := []string{serviceName}
	ip := []string{getIp()}
	port := []string{port}

	m["namespaceId"] = namespaceId
	m["serviceName"] = serviceName
	m["ip"] = ip
	m["port"] = port

	err := postFrom(URL+"/service/create", m)

	if err != nil {
		ServiceCreate = false
		go retryCreateService()
	}

	ServiceCreate = true

	s := new(Service)
	s.NamespaceId = namespaceId[0]
	s.ServiceName = serviceName[0]
	s.Port = port[0]
	s.Ip = ip[0]

	service = s
	fmt.Println(" create service suc")

}

func HealthCheck() {

	for {

		time.Sleep(20 * time.Second)
		header := http.Header{}

		m := make(map[string]string)
		m["namespaceId"] = "SPACE"
		m["serviceName"] = "LOCAL_SERVICE"

		get(URL+"service/heartbeat", header, 1000, m)

		fmt.Println("check health")
	}

}

func retryCreateService() {

	if ServiceCreate {
		return
	}

	time.Sleep(15 * time.Second)

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
