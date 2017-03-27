package db

import (
	"github.com/ps_web_0/models"
	"github.com/vattle/sqlboiler/queries"
	"time"
)

type UserResume struct {
	models.User `boil:",bind"`
	models.ResumeReview   `boil:",bind"`
	models.Resume `boil:",bind"`
}
func (resume *UserResume) FormattedDate() string {
	return resume.Resume.LastUpdatedAt.Format(time.Stamp)
}
func GetModeratorResumes(moderatorId int) ([]*UserResume) {
	var userResume []*UserResume

	queries.RawG("SELECT * FROM resume_review " +
		"JOIN Resume r ON resume_review.resume_id = r.resume_id " +
		"JOIN Users u ON u.user_id = r.author_id " +
		"WHERE resume_review.moderator_id = $1 " +
		" ORDER BY u.Name, r.last_updated_at DESC", moderatorId).BindP(&userResume)

	return userResume
}