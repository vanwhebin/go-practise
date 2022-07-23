package main

import (
	"go-practise/chapt09/controller"
	"go-practise/chapt09/model"
	"net/http"

	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// set DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Setup Controller
	controller.Startup()
	http.ListenAndServe(":8888", context.ClearHandler(http.DefaultServeMux))
}
