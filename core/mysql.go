package core

import (
	"github.com/jinzhu/gorm"
	"log"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
