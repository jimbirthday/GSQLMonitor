package domain

//服务
type Service struct {
	Name        string
	Token       string              //服务授权码
	CheckNums   int                 //检查次数
	NamespaceId string              //命名空间
	ClustetMap  map[string]*Cluster //实例簇
}

//实例簇
type Cluster struct {
	ClustetMap map[string]*Instance //实例
}

//实例
type Instance struct {
	Ip          string
	Port        int
	ServiceName string
	ClusterName string
}
