package application_controller

import (
	"net/http"
	"t-card/dtos/requests"
	"t-card/models"
	"t-card/repository"

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

	existApp, err := repository.FindApplicationByJobIDAndUserID(appReq.JobID, *user_data.ID)

	if err != nil {
		println(err)
	}

	if existApp.ID != nil {
		ctx.AbortWithStatusJSON(http.StatusInsufficientStorage, gin.H{
			"message": "you already applicated.",
		})
		return
	}

	appr, errDb := repository.StoreApplication(*app)
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot create data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data saved successfully.",
		"data":    appr,
	})
}
