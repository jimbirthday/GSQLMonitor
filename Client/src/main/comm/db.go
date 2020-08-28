package comm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	DBUser = "root"
	DBPwd  = "jimbir8520"
	DBUrl  = "localhost"
	DBDatabase = "prison_public"
	//DBPwd      = "Lecent@123"
	//DBUrl      = "192.168.1.15"
	//DBDatabase = "lecent_park_admin"
	DBType     = "mysql"
	PORT       = 3306
)

func InitDB() {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		DBUser,
		DBPwd,
		DBUrl,
		PORT,
		DBDatabase)
	DBINFOURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		DBUser,
		DBPwd,
		DBUrl,
		PORT,
		"information_schema")
	db, err := xorm.NewEngine(DBType, url)
	db1, err := xorm.NewEngine(DBType, DBINFOURL)
	if err != nil {
		panic("db init failed " + err.Error())
	}
	Db = db
	DbInfo = db1
}
