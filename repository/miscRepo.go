package repository

import (
	"t-card/database"
	"t-card/models"
)

func PostFeedback(feedback *models.Feedback) (models.Feedback, error) {
	errDb := database.DB.Table("feedbacks").Create(&feedback).Error
	return *feedback, errDb
}

func GetAllFeedback() ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	errDb := database.DB.Table("feedbacks").Preload("Feeder").Find(&feedbacks).Error
	return feedbacks, errDb
}
