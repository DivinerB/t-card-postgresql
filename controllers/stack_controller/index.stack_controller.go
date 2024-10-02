package stack_controller

import (
	"net/http"
	"t-card/database"
	"t-card/dtos/requests"
	"t-card/models"

	"github.com/gin-gonic/gin"
)

func StoreStack(ctx *gin.Context) {
	var stackReq requests.StackRequest
	if errReq := ctx.ShouldBindJSON(&stackReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	stack := new(models.Stack)
	stack.Name = &stackReq.Name
	stack.Logo = &stackReq.Logo

	errDb := database.DB.Table("job_stacks").Create(&stack).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot create data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data saved successfully.",
		"data":    stack,
	})
}
