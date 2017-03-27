package controller

import (
	"net/http"
	"ps_web_0/view"
	"ps_web_0/db"
)

//Update to logged in user
const moderatorId  = 6

func moderatorResumeSummary(w http.ResponseWriter, r *http.Request)  {
	var data = make(map[string]interface{})


	//db.GetModeratorResumes()
	data["notification"] = "Resumes"

	resumes := db.GetModeratorResumes(moderatorId)
	var splitResumes = make(map[float64][]*db.UserResume)

	for _, element := range resumes {
		splitResumes[element.UserID] = append(splitResumes[element.UserID], element)
	}

	data["user_resumes"] = splitResumes

	//db.GetModeratorResumes(moderatorId)[0].ResumeReview.Review
	view.RenderTemplate(w, "moderator_summary", data)
}