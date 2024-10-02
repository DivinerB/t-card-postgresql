package repository

import (
	"t-card/database"
	"t-card/models"
)

func FindApplicationByJobIDAndUserID(jobID uint, userID uint) (models.Application, error) {
	var app models.Application
	errDB := database.DB.Table("applications").Where("job_id=?", jobID).Where("user_id=?", userID).First(&app).Error
	return app, errDB
}

func StoreApplication(app models.Application) (models.Application, error) {
	err := database.DB.Table("applications").Create(&app).Error
	return app, err
}
