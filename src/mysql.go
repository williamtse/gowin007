package src

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	sqldriver  = "mysql"
	dbhostsip  = "127.0.0.1:3306" //IP地址
	dbusername = "root"           //用户名
	dbpassword = ""               //密码
	dbname     = "qqtydev"        //

)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open(sqldriver, dbusername+":"+dbpassword+"@tcp("+dbhostsip+")/"+dbname+"?charset=utf8")

	return db, err
}

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
