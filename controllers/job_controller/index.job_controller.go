package job_controller

import (
	"net/http"
	"strconv"
	"t-card/database"
	"t-card/dtos/requests"
	"t-card/models"
	"t-card/repository"

	"github.com/gin-gonic/gin"
)

func StoreJob(ctx *gin.Context) {
	var jobReq requests.JobRequest
	if errReq := ctx.ShouldBindJSON(&jobReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}
	user, _ := ctx.Get("user")
	user_data := user.(models.User)

	job := new(models.Job)
	job.Title = &jobReq.Title
	job.Description = &jobReq.Description
	job.Salary = &jobReq.Salary
	job.Money = &jobReq.Money
	job.Geoposition = &jobReq.Geoposition
	job.Owner = &user_data

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

func GetAllJobsWithStacks(ctx *gin.Context) {
	var jobs []models.Job
	errDb := database.DB.Table("jobs").Preload("Applications").Preload("Applications.Applicant").Find(&jobs).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot load data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data loaded successfully.",
		"data":    jobs,
	})
}

func SetBookmark(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userData := user.(models.User)
	userID := *userData.ID

	IDStr := ctx.Param("id")
	ID64, err := strconv.ParseUint(IDStr, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "param invalid",
		})
	}
	jobID := uint(ID64)

	err = repository.SetBookmark(userID, jobID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot load data.",
			"error":   &err,
		})
		return
	}
	bookmarks, errGet := repository.GetBookmarksGroupedByUserID(userID)
	if errGet != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot load data.",
			"error":   &errGet,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "operated successfully",
		"data":    &bookmarks,
	})
}
