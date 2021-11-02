package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("Hello world!"))
		fmt.Fprintf(w, "Hello world!")
	})
	http.ListenAndServe(":8888", nil)
}
