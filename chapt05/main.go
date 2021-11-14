package main

import (
	"go-practise/chapt05/controller"
	"net/http"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":8888", nil)
}
