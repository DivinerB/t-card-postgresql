package repository

import (
	"t-card/database"
	"t-card/models"
)

func GetJobWithApplications(jobs []models.Job) ([]models.Job, error) {
	errDb := database.DB.Table("jobs").Preload("Applications").Find(&jobs).Error
	return jobs, errDb
}
