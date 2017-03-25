package controller

import (
	"net/http"
	"ps_web_0/view"
)

func AdminGET(w http.ResponseWriter, _ *http.Request)  {
	var data map[string]interface{}

	data["text"] = "Hello World"
	view.RenderTemplate(w, "admin", data)
}