package common

import (
	"backEnd_Gin/model"
	"gorm.io/gorm/schema"

	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

//默认连接方式
//var db, err = sql.Open("mysql", dsn)

// 使用gorm连接数据库

var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},
})

func InitDB() *gorm.DB {
	if err != nil {
		fmt.Println("mysql connect failed, detail is [%v]", err.Error())
	} else {
		fmt.Println("use catering database connect success")
	}
	Migrate(db)
	return db
}

// *sql.DB
func GetDB() *gorm.DB {
	return db
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.Student{}, &model.User{})

	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetError() error {
	return err
}
