package model

import (
	"go-practise/chapt07/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func ConnectToDB() *gorm.DB {
	connectingStr := config.GetMysqlConnectingString()
	db, err := gorm.Open("mysql", connectingStr)
	if err != nil {
		panic("Fail to connenct database")
	}

	db.SingularTable(true)
	return db

}
