package initial

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitSql() {
	user := beego.AppConfig.String("mysqluser")
	passwd := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlhost")
	port, err := beego.AppConfig.Int("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")

	if err != nil {
		port = 3306
	}

	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4", user, passwd, host, port, dbname))
}
