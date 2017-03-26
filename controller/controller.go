package controller

import (
	"fmt"
	"net/http"

	"github.com/apexskier/httpauth"
	"ps_web_0/db"
	"ps_web_0/view"
	"github.com/gorilla/mux"
)

type UserForm struct {
	Username string
	Password string
}

var (
	authHandler httpauth.Authorizer
)

func registerRoutesToFuncs(r *mux.Router) {
	r.HandleFunc("/", homeHandler)

	r.HandleFunc("/admin/roles", adminRoles).Methods("GET", "POST")
	r.HandleFunc("/admin/addjob", adminAddJob).Methods("GET", "POST")
	r.HandleFunc("/admin/companies", adminCompanies).Methods("GET", "POST")

	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
}

// Handles the index page, renders a home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("homeHandler Executing...")
	data := map[string]interface{}{
		"hello-user":  103,
		"thats-rught": 104,
	}

	dbAuth, err := httpauth.NewSqlAuthBackend("postgres", db.PsqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	roles := make(map[string]httpauth.Role)
	roles["user"] = 30
	roles["admin"] = 80

	authHandler, err = httpauth.NewAuthorizer(dbAuth, []byte("cookie-encryption-key"), "user", roles)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	username := "roko"
	defaultUser := httpauth.UserData{Username: username, Role: "admin"}
	err = dbAuth.SaveUser(defaultUser)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	// Update user with a password and email address
	err = authHandler.Update(w, r, username, "adminadmin", "admin@localhost.com")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	view.RenderTemplate(w, "index", data)
}

// loginHandler
func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginHandler Executing...")
	data := map[string]interface{}{
		"login": "herro",
		"lol":   "103",
	}

	view.RenderTemplate(w, "login", data)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
	}

	username := r.FormValue("lg_username")
	password := r.FormValue("lg_password")

	if err = authHandler.Login(w, r, username, password, "/"); err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	fmt.Println("ff")
}

func Init() {
	router := mux.NewRouter()

	registerRoutesToFuncs(router)

	http.Handle("/", router)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	fmt.Println("Public directory launched;")
}
