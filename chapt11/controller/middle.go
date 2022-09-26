package controller

import (
	"go-practise/chapt11/model"
	"log"
	"net/http"
)

func middleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := getSessionUser(r)
		log.Println("middle: ", username)
		log.Println("middle error: ", err)

		if username != "" {
			log.Println("Last seen:", username)
			model.UpdateLastSeen(username)
		}

		if err != nil {
			log.Println("middle get Session err and redirect to login")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
