package repository

import (
	"t-card/database"
	"t-card/models"
)

func GetJobWithApplications(jobs []models.Job) ([]models.Job, error) {
	errDb := database.DB.Table("jobs").Preload("Applications").Find(&jobs).Error
	return jobs, errDb
}

func SetBookmark(userID uint, jobID uint) error {
	// errDb := database.DB.Table("bookmarks").Where("user_id = ?", userID).Where("job_id = ?", jobID)
	var bookmark models.Bookmark
	bookmark.JobID = &jobID
	bookmark.UserID = &userID
	errDB := database.DB.Table("bookmarks").Where(&bookmark).First(&bookmark).Error
	println(errDB)
	if bookmark.ID != nil {
		errDB = database.DB.Table("bookmarks").Delete(&bookmark).Error
		return errDB
	}
	errDB = database.DB.Table("bookmarks").Create(&bookmark).Error
	return errDB
}

func GetBookmarksGroupedByUserID(userID uint) ([]models.Bookmark, error) {
	var bookmarks []models.Bookmark
	errDB := database.DB.Table("bookmarks").Where("user_id = ?", userID).Find(&bookmarks).Error
	return bookmarks, errDB
}
