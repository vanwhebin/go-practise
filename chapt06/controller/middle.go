package controller

import (
	"net/http"
)

func middleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := getSessionUser(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		} else {
			if r.RequestURI == "/login" {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			}

			next.ServeHTTP(w, r)
		}
	}
}
