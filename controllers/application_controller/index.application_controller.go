package application_controller

import (
	"net/http"
	"t-card/database"
	"t-card/dtos/requests"
	"t-card/models"

	"github.com/gin-gonic/gin"
)

func StoreApplication(ctx *gin.Context) {
	var appReq requests.ApplicantionRequest
	if errReq := ctx.ShouldBindJSON(&appReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}
	user, _ := ctx.Get("user")
	user_data := user.(models.User)

	app := new(models.Application)
	app.JobID = &appReq.JobID
	app.Applicant = &user_data

	errDb := database.DB.Table("applications").Create(&app).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot create data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data saved successfully.",
		"data":    app,
	})
}
