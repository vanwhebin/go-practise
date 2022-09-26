package controller

import (
	"go-practise/chapt07/model"
	"net/http"
)

func middleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := getSessionUser(r)
		if username != "" {
			model.UpdateLastSeen(username)
		}

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
