package model

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

type Database struct {
	Self *gorm.DB
}

var Db *Database

func (db *Database) DbInit() {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))

	newDb, err := gorm.Open("mysql",dns)
	if err != nil {
		log.Println(err)
	}

	//设置表名默认为结构体首字母小写形式
	newDb.SingularTable(true)
	//开启gorm的logger
	newDb.LogMode(true)

	Db = &Database{Self: newDb}
}

func (db *Database) DbClose() {
	if err := Db.Self.Close(); err != nil {
		log.Println(err)
	}
	return
}


