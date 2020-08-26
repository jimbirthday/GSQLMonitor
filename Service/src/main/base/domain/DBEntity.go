package domain

//服务
type Service struct {
	Name        string
	Token       string    //服务授权码
	CheckTime   int64     //检查时间
	CheckNums   int       //检查次数
	NamespaceId string    //命名空间
	Ins         *Instance //实例簇
}

//实例
type Instance struct {
	Ip   string
	Port int
}
