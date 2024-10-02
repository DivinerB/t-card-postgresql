package repository

import (
	"t-card/database"
	"t-card/models"
)

func GetUserByTID(id string) (models.User, error) {
	var user models.User
	errDB := database.DB.Table("users").Where("t_id=?", id).First(&user).Error

	return user, errDB
}
