package controller

import (
	"html/template"

	"go-practise/chapt08/config"

	"github.com/gorilla/sessions"
)

var (
	homeController home
	templates      map[string]*template.Template
	sessionName    string
	sessionStore   *sessions.CookieStore
)

func init() {
	templates = PopulateTemplates()
	sessionConfig := config.GetSessionConfig()
	sessionStore = sessions.NewCookieStore([]byte(sessionConfig.AuthKey), []byte(sessionConfig.EncryptKey))
	sessionStore.Options = &sessions.Options{
		HttpOnly: sessionConfig.HttpOnly,
		MaxAge:   sessionConfig.MaxAge,
	}
	sessionName = sessionConfig.Name
}

// Startup func
func Startup() {
	homeController.registerRoutes()
}
