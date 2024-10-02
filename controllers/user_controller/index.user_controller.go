package user_controller

import (
	"net/http"
	"strconv"
	"t-card/database"
	"t-card/dtos/requests"
	"t-card/dtos/responses"
	"t-card/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUser(ctx *gin.Context) {

	users := new([]models.User)

	if err := database.DB.Table("users").Find(&users).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetUserByID(ctx *gin.Context) {

	id := ctx.Param("id")
	user := new(responses.UserResponse)

	err := database.DB.Table("users").Where("id=?", id).Find(&user).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error.",
		})
		return
	}

	if user.Name == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data transmitted.",
		"data":    user,
	})

}

func StoreUser(ctx *gin.Context) {
	var userReq requests.UserRequest

	if errReq := ctx.ShouldBindJSON(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})

		return
	}

	userAlreadyExist := new(models.User)

	database.DB.Table("users").Where("t_id=?", userReq.TID).First(&userAlreadyExist)

	if userAlreadyExist.TID != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "telegram id already used.",
		})

		return
	}

	user := new(models.User)

	hash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot hash password.",
		})
		return
	}
	hashStr := string(hash)

	user.Name = &userReq.Name
	user.TID = &userReq.TID
	user.Password = &hashStr
	user.Role = &userReq.Role
	user.Birthday = &userReq.Birthday

	errDb := database.DB.Table("users").Create(&user).Error

	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot create data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data saved successfully.",
		"data":    user,
	})
}

func UpdateUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	user := new(models.User)

	userReq := new(requests.UserRequest)

	userAlreadyExist := new(models.User)

	if errReq := ctx.ShouldBindJSON(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})

		return
	}

	errDb := database.DB.Table("users").Where("id=?", id).Find(&user).Error

	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error.",
		})

		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found.",
		})
		return
	}

	errUserAlreadyExist := database.DB.Table("users").Where("tid=?", userReq.TID).Find(&userAlreadyExist).Error

	if errUserAlreadyExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error.",
		})

		return
	}

	if userAlreadyExist.TID != nil && *user.ID != *userAlreadyExist.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "TID already used.",
		})
		return
	}

	user.Name = &userReq.Name
	user.TID = &userReq.TID
	user.Password = &userReq.Password
	user.Role = &userReq.Role
	user.Birthday = &userReq.Birthday

	errUpdate := database.DB.Table("users").Where("id=?", id).Updates(&user).Error

	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't update data.",
		})
		return
	}

	userResponse := responses.UserResponse{
		Name:     user.Name,
		TID:      user.TID,
		Password: user.Password,
		Role:     user.Role,
		Birthday: user.Birthday,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "data updated successfully.",
		"data":    userResponse,
	})
}

func DeleteUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	user := new(models.User)

	errFind := database.DB.Table("users").Where("id=?", id).Find(&user).Error

	if errFind != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error.",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found.",
		})
		return
	}

	errDb := database.DB.Table("users").Unscoped().Where("id=?", id).Delete(&models.User{}).Error

	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error.",
			"error":   errDb.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "data deleted successfully.",
	})
}

// Paginate
func GetUserPaginate(ctx *gin.Context) {
	page := ctx.Query("page")

	if page == "" {
		page = "1"
	}

	perPage := ctx.Query("per_page")

	if perPage == "" {
		perPage = "10"
	}

	pageInt, _ := strconv.Atoi(page)
	perPageInt, _ := strconv.Atoi(perPage)

	if pageInt < 1 {
		pageInt = 1
	}

	users := new([]models.User)

	if err := database.DB.Table("users").Offset((pageInt - 1) * perPageInt).Limit(perPageInt).Find(&users).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":     users,
		"page":     pageInt,
		"per_page": perPageInt,
	})
}
