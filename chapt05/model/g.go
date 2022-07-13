package model

import (
	"fmt"
	"go-practise/chapt05/config"
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func ConnectToDB() *gorm.DB {
	connectingStr := config.GetMysqlConnectingString()
	log.Println("Connet to db...")
	log.Println(connectingStr)
	db, err := gorm.Open("mysql", connectingStr)
	fmt.Println(connectingStr)
	if err != nil {
		panic("Fail to connenct database")
	}

	db.SingularTable(true)
	return db

}
