package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUsername string = "root"
	dbPassword string = "123456"
	ipAddress  string = "127.0.0.1"
	port       int    = 3306
	dbName     string = "student"
	charset    string = "utf8"
)

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", dbUsername, dbPassword, ipAddress, port, dbName, charset)

var db, err = sql.Open("mysql", dsn)

func InitDB() *sql.DB {
	if err != nil {
		fmt.Println("mysql connect failed, detail is [%v]", err.Error())
	} else {
		fmt.Println("use catering database connect success")
	}
	return db
}

func GetDB() *sql.DB {
	return db
}

func GetError() error {
	return err
}
