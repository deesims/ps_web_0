package controller

import (
	"net/http"
	"github.com/ps_web_0/view"
	"github.com/ps_web_0/db"
	"github.com/ps_web_0/models"
	"strconv"
)

//Update to logged in user
const moderatorId  = 6

func moderatorResumeSummary(w http.ResponseWriter, r *http.Request)  {
	var data = make(map[string]interface{})

	if r.Method == "POST" {
		resumeId, _ := strconv.ParseFloat(r.PostFormValue("ResumeID"), 64)
		moderatorId, _ := strconv.ParseFloat(r.PostFormValue("ModeratorID"), 64)

		review := r.PostFormValue("Review")

		resumeReview := models.FindResumeReviewGP(resumeId, moderatorId)
		resumeReview.Review = review
		resumeReview.UpdateGP()

		data["notification"] = "Updated"
	}

	resumes := db.GetModeratorResumes(moderatorId)
	var splitResumes = make(map[float64][]*db.UserResume)

	for _, element := range resumes {
		splitResumes[element.UserID] = append(splitResumes[element.UserID], element)
	}

	data["users_resumes"] = splitResumes

	view.RenderTemplate(w, "moderator_summary", data)
}