package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/sir")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
