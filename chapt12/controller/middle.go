package controller

import (
	"go-practise/chapt12/model"
	"net/http"
)

func middleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := getSessionUser(r)

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		if r.RequestURI == "/login" {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		if r.RequestURI == "/register" {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		if username != "" {
			model.UpdateLastSeen(username)
		}

		next.ServeHTTP(w, r)
	}
}
