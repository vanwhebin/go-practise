package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Username string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := User{Username: "wanweibin"}
		tpl, _ := template.ParseFiles("../templates/index.html")
		tpl.Execute(w, &user)
	})
	http.ListenAndServe(":8888", nil)
}
