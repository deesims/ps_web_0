package view

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	templates     *template.Template
	templateError error
)

func init() {
	templates, templateError = template.ParseFiles("view/login.html", "view/index.html", "view/admin.html")
	if templateError != nil {
		panic(templateError)
	}
}
func RenderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		fmt.Println("Executing template failed...")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
