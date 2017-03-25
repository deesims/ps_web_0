package view

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates, _ = template.ParseFiles("view/login.html", "view/index.html", "view/admin.html") // add to this list to render templates

func RenderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {

	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		fmt.Println("Executing template failed...")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
