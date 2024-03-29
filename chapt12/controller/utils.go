package controller

import (
	"crypto/tls"
	"errors"
	"fmt"
	"go-practise/chapt12/config"
	"go-practise/chapt12/vm"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"gopkg.in/gomail.v2"
)

// PopulateTemplates func
// Create map template name to template.Template
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

func getSessionUser(r *http.Request) (string, error) {
	var username string
	session, err := sessionStore.Get(r, sessionName)
	if err != nil {
		return "", err
	}

	val := session.Values["user"]
	username, ok := val.(string)
	if !ok {
		return "", errors.New("can not get session user")
	}
	return username, nil
}

func setSessionUser(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := sessionStore.Get(r, sessionName)

	if err != nil {
		return err
	}
	session.Values["user"] = username
	session.Values["authenticated"] = true
	err = session.Save(r, w)

	if err != nil {
		return err
	}
	return nil
}

func clearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := sessionStore.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	session.Values["authenticated"] = false
	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func checkLen(fieldName, fieldValue string, minLen, maxLen int) string {
	lenField := len(fieldValue)
	if lenField < minLen {
		return fmt.Sprintf("%s field is too short, less than %d", fieldName, minLen)
	}
	if lenField > maxLen {
		return fmt.Sprintf("%s field is too long, more than %d", fieldName, maxLen)
	}
	return ""
}

func checkUsername(username string) string {
	return checkLen("Username", username, 3, 20)
}

func checkPassword(password string) string {
	return checkLen("Password", password, 6, 50)
}

func checkEmail(email string) string {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return "invvalid email"
	}
	return ""
}

func checkUserPassword(username, password string) string {
	if !vm.CheckLogin(username, password) {
		return "Username and password is not correct."
	}
	return ""
}

func checkUserExist(username string) string {
	if !vm.CheckUserExist(username) {
		return fmt.Sprintf("%s already exist, please choose another username", username)
	}
	return ""
}

// checkLogin()
func checkLogin(username, password string) []string {
	var errs []string
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkPassword(password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkUserPassword(username, password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

// checkRegister()
func checkRegister(username, email, pwd1, pwd2 string) []string {
	var errs []string
	if pwd1 != pwd2 {
		errs = append(errs, "2 password does not match")
	}
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkPassword(pwd1); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkEmail(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkUserExist(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

// addUser()
func addUser(username, password, email string) error {
	return vm.AddUser(username, password, email)
}

func setFlash(w http.ResponseWriter, r *http.Request, message string) {
	session, _ := sessionStore.Get(r, sessionName)
	session.AddFlash(message, flashName)
	session.Save(r, w)
}

func getFlash(w http.ResponseWriter, r *http.Request) string {
	session, _ := sessionStore.Get(r, sessionName)
	fm := session.Flashes(flashName)
	if fm == nil {
		return ""
	}
	session.Save(r, w)
	return fmt.Sprintf("%v", fm[0])

}

func getPage(r *http.Request) int {
	url := r.URL
	query := url.Query()

	p := query.Get("page")
	if p == "" {
		return 1
	}

	page, err := strconv.Atoi(p)
	if err != nil {
		return 1
	}

	return page
}

func sendEmail(target, subject, content string) {
	emailConfig := config.GetEmailConfig()
	d := gomail.NewDialer(emailConfig.Smtp, emailConfig.Port, emailConfig.User, emailConfig.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", emailConfig.User)
	m.SetHeader("To", target)
	m.SetAddressHeader("Cc", emailConfig.User, "admin")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Email Error:", err)
		return
	}
}
