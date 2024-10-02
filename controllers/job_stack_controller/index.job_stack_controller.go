package jobstack_controller

import (
	"net/http"
	"t-card/database"
	"t-card/dtos/requests"
	"t-card/models"

	"github.com/gin-gonic/gin"
)

func StoreJobStack(ctx *gin.Context) {
	var jobStackReq requests.JobStackRequest
	if errReq := ctx.ShouldBindJSON(&jobStackReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	jobStack := new(models.Job_Stack)
	jobStack.Name = &jobStackReq.Name
	jobStack.Logo = &jobStackReq.Logo

	errDb := database.DB.Table("job_stacks").Create(&jobStack).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot create data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data saved successfully.",
		"data":    jobStack,
	})
}
