package auth_contoller

import (
	"net/http"
	"t-card/database"
	"t-card/models"
	"t-card/requests"
	"t-card/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *gin.Context) {

	loginReq := new(requests.LoginRequest)

	err := ctx.ShouldBindJSON(&loginReq)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user := new(models.User)

	errUser := database.DB.Table("users").Where("t_id = ?", loginReq.TID).Find(&user).Error

	if errUser != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "credential not valid.",
		})
		return
	}

	if user.TID == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "credential not valid.",
		})
		return
	}

	// Check Password
	if loginReq.Password != *user.Password {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "credential not valid.",
		})
		return
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"tid":  user.TID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token, errToken := utils.GenerateToken(&claims)

	if errToken != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed generate token.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login succesfully",
		"token":   token,
	})
}
