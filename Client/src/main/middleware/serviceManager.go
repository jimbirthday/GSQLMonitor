package middleware

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type Service struct {
	NamespaceId string `yaml:"namespaceId"`
	ServiceName string `yaml:"serviceName"`
	GroupName   string `yaml:"groupName"`
	Ip          string
	Port        string `yaml:"port"`
}

var (
	service       *Service
	ServiceCreate bool
	URL           string
	namespaceId   string
	serviceName   string
	port          string
)

func init() {
	fileReadConf()
	go CreateService()
	go HealthCheck()
}

type conf struct {
	URL         string `yaml:"serviceUrl"`
	NamespaceId string `yaml:"namespace"`
	ServiceName string `yaml:"name"`
	Port        string `yaml:"port"`
}

func fileReadConf() {

	yamlFile, err := ioutil.ReadFile("src/main/config/config.yaml")

	if err != nil {
		log.Fatal("conf err " + err.Error())
	}

	c := new(conf)
	err = yaml.Unmarshal(yamlFile, c)
	URL = c.URL
	namespaceId = c.NamespaceId
	serviceName = c.ServiceName
	port = c.Port
	if err != nil {
		log.Fatal("conf err " + err.Error())
	}

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
		m["namespaceId"] = namespaceId
		m["serviceName"] = serviceName

		get(URL+"/service/heartbeat", header, 1000, m)

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
