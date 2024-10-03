package misc_controller

import (
	"net/http"
	"t-card/dtos/requests"
	"t-card/models"
	"t-card/repository"

	"github.com/gin-gonic/gin"
)

func PostFeedback(ctx *gin.Context) {
	var feedbackReq requests.FeedbackRequest
	if errReq := ctx.ShouldBindJSON(&feedbackReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}
	user, _ := ctx.Get("user")
	userData := user.(models.User)

	feedback := new(models.Feedback)
	feedback.Title = &feedbackReq.Title
	feedback.Description = &feedbackReq.Description
	feedback.Feeder = &userData
	feedbackRes, err := repository.PostFeedback(feedback)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot create data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data saved successfully.",
		"data":    feedbackRes,
	})
}

func GetAllFeedbackWithUser(ctx *gin.Context) {
	feedbacks, errDb := repository.GetAllFeedback()
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot load data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data loaded successfully.",
		"data":    feedbacks,
	})
}
