package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type User struct {
	Username string
}

type Post struct {
	User
	Body string
}

type IndexViewModel struct {
	Title string
	User
	Posts []Post
}

func PopulateTemplates() map[string]*template.Template {
	const basePath = "../templates"
	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(basePath + "/_base.html"))

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user1 := User{Username: "wanweibin"}
		user2 := User{Username: "wwb"}

		posts := []Post{
			{User: user1, Body: "beautiful day in China"},
			{User: user2, Body: "beautiful day in SZ"},
		}

		v := IndexViewModel{Title: "TEST HOMEPAGE", User: user1, Posts: posts}

		tpl := PopulateTemplates()
		tpl["index.html"].Execute(w, &v)
	})
	http.ListenAndServe(":8888", nil)
}
