package main

import (
	"go-practise/chapt04/controller"
	"net/http"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":8888", nil)
}
