package auth_contoller

import (
	"net/http"
	"t-card/database"
	"t-card/dtos/requests"
	"t-card/models"
	"t-card/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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
			"message": "user not valid.",
		})
		return
	}

	// Compare the password
	errPsw := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(loginReq.Password))

	if errPsw != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "password not valid",
		})
		return
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token, errToken := utils.GenerateToken(&claims)

	if errToken != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed generate token.",
		})
		return
	}

	// set cookie
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	// send it as a response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
		"token":   token,
	})
}
