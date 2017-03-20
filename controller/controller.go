package controller

import (
	"fmt"
	"net/http"

	"github.com/deesims/ps_web_0/db"
	"github.com/deesims/ps_web_0/models"
	"github.com/deesims/ps_web_0/view"
	"github.com/gorilla/mux"
)

type UserForm struct {
	Username string
	Password string
}

func registerRoutesToFuncs(r *mux.Router) {
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/login", loginPostHandler).Methods("GET")
	r.HandleFunc("/login", loginGetHandler).Methods("POST")
}

// Handles the index page, renders a home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("homeHandler Executing...")
	data := map[string]interface{}{
		"hello-user":  103,
		"thats-rught": 104,
	}
	view.RenderTemplate(w, "index", data)
}

// loginHandler
func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginHandler Executing...")
	data := map[string]interface{}{
		"login": "herro",
		"lol":   "103",
	}

	view.RenderTemplate(w, "login", data)
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
	}

	UserForm := &UserForm{
		r.FormValue("lg_username"),
		r.FormValue("lg_password"),
	}

	db := db.Connection()
	defer db.Close()

	user, err := models.Users(db).All()
	if err != nil {
		fmt.Println(err)
	}
	// elided error check

	fmt.Println("form n others", UserForm)
	fmt.Println("first user of system", user)
}

func Init() {
	router := mux.NewRouter()

	registerRoutesToFuncs(router)

	http.Handle("/", router)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	fmt.Println("Public directory launched;")
}
