package repository

import (
	"t-card/database"
	"t-card/models"
)

func GetApplicationByJobIDAndUserID(jobID uint, userID uint) (models.Application, error) {
	var app models.Application
	errDB := database.DB.Table("applications").Where("job_id=?", jobID).Where("user_id=?", userID).First(&app).Error
	return app, errDB
}

func StoreApplication(app models.Application) (models.Application, error) {
	err := database.DB.Table("applications").Create(&app).Error
	return app, err
}

func UpdateApplicationState(id uint, state string) (models.Application, error) {
	var app models.Application
	err := database.DB.Where("id=?", id).First(&app).Update("state", state).Error
	return app, err
}
