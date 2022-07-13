package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Username string
}

type IndexViewModel struct {
	Title string
	User  User
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := User{Username: "wanweibin"}
		v := IndexViewModel{Title: "TEST HOMEPAGE", User: user}
		tpl, _ := template.ParseFiles("../../templates/blog.html")
		tpl.Execute(w, &v)
	})
	http.ListenAndServe(":8888", nil)
}
