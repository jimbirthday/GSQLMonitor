package domain

type ProcesslistEntity struct {
	Id       string
	BaseName string
	User     string
	Command  string
	Time     string
	State    string
	Info     string
	Host     string
}

type QuestionsEntity struct {
	VariableName string
	Value        float64
}

type KeyBufferEntity struct {
	Name 		 string
	Value        float64
}

type DBInfo struct {
	Port int
	DataBaseName string
	Engine string
	DataSize string
	Uptime string
	Status string
	SlowQueryLog string
	SlowQueryLogFile string
	LongQueryTime string
}

type SlowEntity struct {
	Path string
	Counts string
	FileLog string

}