package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/deesims/ps_web_0/db"
	"github.com/gorilla/mux"
)

var templates, _ = template.ParseFiles("login.html", "index.html")

// Handles the index page, renders a home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("homeHandler Executing...")
	data := map[string]interface{}{
		"hello-user":  103,
		"thats-rught": 104,
	}
	renderTemplate(w, "index", data)
}

// loginHandler
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginHandler Executing...")
	data := map[string]interface{}{
		"login": "herro",
		"lol":   "103",
	}
	renderTemplate(w, "login", data)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {

	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		fmt.Println("Executing template failed...")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/login", loginHandler)

	http.Handle("/", r)
	//http.Handle("/login", r)

	db.InsertUser()

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Public directory launched;")

	http.ListenAndServe("localhost:7388", nil)

}
