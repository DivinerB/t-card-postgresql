package application_controller

import (
	"fmt"
	"net/http"
	"strconv"
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

	existApp, err := repository.GetApplicationByJobIDAndUserID(appReq.JobID, *user_data.ID)

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

func UpdateApplication(ctx *gin.Context) {
	var appUpdateReq requests.ApplicantionUpdateRequest
	if errReq := ctx.ShouldBind(&appUpdateReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}
	appIDStr := ctx.Param("id")
	appID64, err := strconv.ParseUint(appIDStr, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "param invalid",
		})
	}
	appID := uint(appID64)
	app, err := repository.UpdateApplicationState(appID, appUpdateReq.State)
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "data changed successfully.",
		"data":    &app,
	})
}
