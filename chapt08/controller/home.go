package controller

import (
	"fmt"
	"go-practise/chapt08/vm"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type home struct{}

func (h home) registerRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", middleAuth(indexHandler))
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", middleAuth(logoutHandler))
	r.HandleFunc("/user/{username}", middleAuth(profileHandler))
	r.HandleFunc("/follow/{username}", middleAuth(followHandler))
	r.HandleFunc("/unfollow/{username}", middleAuth(unfollowhandler))
	r.HandleFunc("/profile_edit", middleAuth(profileEditHandler))
	r.HandleFunc("/register", registerHandler)

	http.Handle("/", r)
}

// 首页
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "index.html"
	vop := vm.IndexViewModelOp{}
	username, _ := getSessionUser(r)
	v := vop.GetVM(username)
	templates[tpName].Execute(w, &v)
}

// 登录
func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "login.html"
	vop := vm.LoginViewModelOp{}
	v := vop.GetVM()

	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		errs := checkLogin(username, password)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

// 退出
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w, r)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

// 注册
func registerHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "register.html"
	vop := vm.RegisterViewModelOp{}
	v := vop.GetVM()

	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		email := r.Form.Get("email")
		pwd1 := r.Form.Get("pwd1")
		pwd2 := r.Form.Get("pwd2")

		errs := checkRegister(username, email, pwd1, pwd2)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			if err := addUser(username, pwd1, email); err != nil {
				log.Println("add User error:", err)
				w.Write([]byte("Error insert database"))
				return
			}
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}
	}
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "profile.html"
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)
	vop := vm.ProfileViewModelOp{}
	v, err := vop.GetVM(sUser, pUser)
	if err != nil {
		msg := fmt.Sprintf("user ( %s ) does not exist", pUser)
		w.Write([]byte(msg))
		return
	}
	templates[tpName].Execute(w, &v)
}

func profileEditHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "profile_edit.html"
	username, _ := getSessionUser(r)
	vop := vm.ProfileEditViewModelOp{}
	v := vop.GetVM(username)
	if r.Method == http.MethodGet {
		err := templates[tpName].Execute(w, &v)
		if err != nil {
			log.Println(err)
		}
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		aboutme := r.Form.Get("aboutme")
		log.Println(aboutme)
		if err := vm.UpdateAboutMe(username, aboutme); err != nil {
			log.Println("update Aboutme error:", err)
			w.Write([]byte("Error update aboutme"))
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/user/%s", username), http.StatusTemporaryRedirect)
	}
}

func followHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)
	err := vm.Follow(sUser, pUser)
	if err != nil {
		log.Println("Follow error:", err)
		w.Write([]byte("Error in Follow"))
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/user/%s", pUser), http.StatusTemporaryRedirect)
}

func unfollowhandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)

	err := vm.Unfollow(sUser, pUser)

	if err != nil {
		log.Println("Unfollow error:", err)
		w.Write([]byte("Error in Unfollow"))
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/user/%s", pUser), http.StatusTemporaryRedirect)
}
