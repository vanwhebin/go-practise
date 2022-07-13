package main

import (
	"go-practise/chapt05/controller"
	"go-practise/chapt05/model"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// set DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Setup Controller
	controller.Startup()
	http.ListenAndServe(":8888", nil)
}
