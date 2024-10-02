package routes

import (
	"t-card/controllers/file_controller"
	"t-card/middleware"

	"github.com/gin-gonic/gin"
)

func v1Route(routeGrup *gin.RouterGroup) {

	route := routeGrup

	// ROUTE YANG MEMBUTUHKAN MIDDLEWARE DAPAT DI GRUPKAN
	authRoute := route.Group("file", middleware.RequireAuth)
	authRoute.DELETE("/:filename", file_controller.HandleRemoveFile)
	authRoute.POST("/", file_controller.HandleUploadFile)
	authRoute.POST("/middleware", middleware.RequireAuth, middleware.UploadFile, file_controller.SendStatus)
}
