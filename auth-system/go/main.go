package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/RaiiaaGithub/auth-system/users"
)

func templating(w http.ResponseWriter, fileName string, data interface{}) {
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, data)
}

func signInUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	err := users.DefaultUserService.VerifyUser(newUser)
	if err != nil {
		templating(w, "sign-in.html", "User Sign-in Failure")
		return
	}
	templating(w, "sign-in.html", "User Sign-in Success. Email: "+newUser.Email)
}

func signUpUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	err := users.DefaultUserService.CreateUser(newUser)
	if err != nil {
		templating(w, "sign-up.html", "New User Sign-up Failure")
		return
	}
	templating(w, "sign-up.html", "New User Sign-up Success")
}

func getUser(r *http.Request) users.User {
	email := r.FormValue("email")
	password := r.FormValue("password")
	return users.User{
		Email:    email,
		Password: password,
	}
}

func getSignInPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "sign-in.html", nil)
}

func getSignUpPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "sign-up.html", nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/sign-in":
		signInUser(w, r)
	case "/sign-up":
		signUpUser(w, r)
	case "/sign-in-form":
		getSignInPage(w, r)
	case "/sign-up-form":
		getSignUpPage(w, r)
	}
}

func main() {
	http.HandleFunc("/", userHandler)
	log.Fatal(http.ListenAndServe("", nil))
}
