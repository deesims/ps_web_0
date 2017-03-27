package controller

import (
	"net/http"
	"ps_web_0/view"
)

const moderatorId  = 10

func moderatorResumeSummary(w http.ResponseWriter, r *http.Request)  {
	var data = make(map[string]interface{})


	//db.GetModeratorResumes()
	data["notification"] = "Hmm it works"

	view.RenderTemplate(w, "moderator_summary", data)
}
