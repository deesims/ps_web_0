package controller

import (
	"net/http"
	"ps_web_0/view"
)

func adminGET(w http.ResponseWriter, _ *http.Request) {
	var data = make(map[string]interface{})

	data["text"] = "Hello World"
	view.RenderTemplate(w, "admin", data)
}