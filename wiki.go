// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}


// Handles the index page, renders a home page
func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("homeHandler Executing...")
	renderTemplate(w, "index", "")
}


var templates = template.Must(template.ParseFiles("index.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data string) {

	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		fmt.Println("Executing template failed...")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	fmt.Println("Server Launched...")

	http.Handle("/", r)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.ListenAndServe("localhost:8080", nil)

}
