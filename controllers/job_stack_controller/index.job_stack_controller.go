package jobstackcontroller

import (
	"net/http"
	"t-card/database"
	"t-card/models"
	"t-card/requests"

	"github.com/gin-gonic/gin"
)

func StoreJobStack(ctx *gin.Context) {
	var jobReq requests.JobRequest
	if errReq := ctx.ShouldBindJSON(&jobReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}
	user, _ := ctx.Get("user")
	user_data := user.(models.User)
	println(*user_data.ID)

	job := new(models.Job)
	job.Title = &jobReq.Title
	job.Description = &jobReq.Description
	job.Salary = &jobReq.Salary
	job.Money = &jobReq.Money
	job.Geoposition = &jobReq.Geoposition
	job.UserID = user_data.ID

	errDb := database.DB.Table("jobs").Create(&job).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot create data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data saved successfully.",
		"data":    job,
	})
}
