package controller

import (
	"html/template"

	"go-practise/chapt11/config"

	"github.com/gorilla/sessions"
)

var (
	homeController home
	templates      map[string]*template.Template
	sessionName    string
	sessionStore   *sessions.CookieStore
	flashName      string
	pageLimit      int
)

func init() {
	templates = PopulateTemplates()
	sessionConfig := config.GetSessionConfig()
	basicConfig := config.GetBasicConfig()
	sessionStore = sessions.NewCookieStore([]byte(sessionConfig.AuthKey), []byte(sessionConfig.EncryptKey))
	sessionStore.Options = &sessions.Options{
		HttpOnly: sessionConfig.HttpOnly,
		MaxAge:   sessionConfig.MaxAge,
	}
	sessionName = sessionConfig.Name
	flashName = "go-flash"
	pageLimit = basicConfig.PageLimit
}

// Startup func
func Startup() {
	homeController.registerRoutes()
}
