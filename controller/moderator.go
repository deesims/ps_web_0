package controller

import (
	"fmt"
	"github.com/deesims/ps_web_0/db"
	"github.com/deesims/ps_web_0/models"
	"github.com/deesims/ps_web_0/view"
	"net/http"
	"strconv"
)

func checkModeratorRole(w http.ResponseWriter, r *http.Request) bool {
	currentUser, err := authHandler.CurrentUser(w, r)
	if err != nil {
		fmt.Println("Error getting user: ", err.Error())
		return false
	}
	if currentUser.Role != "moderator" {
		fmt.Println("Current user is not a moderator.")
		return false
	}

	return true
}

func moderatorResumeSummary(w http.ResponseWriter, r *http.Request) {

	if !checkModeratorRole(w, r) {
		fmt.Println("Error, current user is not moderator or not logged in.")
		return
	}

	currentUser, _ := authHandler.CurrentUser(w, r)
	moderator := db.FindUserFromUsername(currentUser.Username)

	var data = make(map[string]interface{})

	if r.Method == "POST" {
		r.ParseForm()
		resumeID, _ := strconv.ParseFloat(r.PostFormValue("ResumeID"), 64)
		moderatorID, _ := strconv.ParseFloat(r.PostFormValue("ModeratorID"), 64)

		review := r.PostFormValue("Review")

		resumeReview := models.FindResumeReviewGP(resumeID, moderatorID)
		resumeReview.Review = review
		resumeReview.UpdateGP()

		data["notification"] = "Updated"
	}

	resumes := db.GetModeratorResumes((int)(moderator.UserID))
	var splitResumes = make(map[float64][]*db.UserResume)

	for _, element := range resumes {
		splitResumes[element.UserID] = append(splitResumes[element.UserID], element)
	}

	data["users_resumes"] = splitResumes

	view.RenderTemplate(w, "moderator_summary", data)
}
