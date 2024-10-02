package routes

import (
	"t-card/config/app_config"
	"t-card/controllers/application_controller"
	"t-card/controllers/book_controller"
	"t-card/controllers/job_controller"
	"t-card/controllers/stack_controller"
	"t-card/controllers/user_controller"
	"t-card/controllers/user_controller/auth_contoller"
	"t-card/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	route := app.Group("api")

	// ROUTE STATIC
	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)

	// ROUTE USER
	userRoute := route.Group("user")
	userRoute.GET("/", user_controller.GetAllUser)
	userRoute.GET("/:id", user_controller.GetUserByID)
	userRoute.GET("/paginate", user_controller.GetUserPaginate)
	userRoute.POST("/", user_controller.StoreUser)
	userRoute.PATCH(":id", user_controller.UpdateUserById)
	userRoute.DELETE("/:id", user_controller.DeleteUserById)

	// ROUTE JOB
	jobRoute := route.Group("job")
	jobRoute.POST("/", middleware.RequireAuth, middleware.RequireEmployer, job_controller.StoreJob)
	jobRoute.GET("/", job_controller.GetAllJobsWithStacks)

	// ROUTE JOB
	stackRoute := route.Group("stack")
	stackRoute.POST("/", middleware.RequireAuth, middleware.RequireFreelancer, stack_controller.StoreStack)
	stackRoute.GET("/", stack_controller.GetAllStacksWithJobs)

	// ROUTE JOB
	applicationRoute := route.Group("application")
	applicationRoute.POST("/", middleware.RequireAuth, middleware.RequireFreelancer, application_controller.StoreApplication)
	// applicationRoute.GET("/", stack_controller.GetAllStacksWithJobs)
	// ROUTE BOOK
	route.GET("/book", book_controller.GetAllBook)

	// ROUTE FILE

	// v1Route(route)

	// LOGIN
	route.POST("/login", auth_contoller.Login)
}
