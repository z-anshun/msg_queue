package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

const (
	USER_NAME = "root"
	PASS_WORD = "123"
	HOST      = "localhost"
	PORT      = "3306"
	DATABASE  = "tickets"
	CHARSET   = "utf8"
)

func InitDB() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
	open, err := gorm.Open("mysql", dbDSN)
	if err != nil {
		log.Panicln("open mysql error")
	}
	e := open.HasTable(&Order{})
	if !e {
		if err := open.CreateTable(&Order{}); err != nil {
			log.Println("creat order error")
		}
	}
	e = open.HasTable(&Tickets{})
	if !e {
		if err := open.CreateTable(&Tickets{}); err != nil {
			log.Println("creat ticket error")
		}
	}

	DB = open

}
