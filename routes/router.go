package routes

import (
	"gin-mvc/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/", controllers.HomeIndex)
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/create", controllers.CreateUser)
		userRoutes.POST("/search", controllers.GetAllUsersByParams)
	}

	adminRoutes := r.Group("/admin")
	{
		adminRoutes.POST("/login", controllers.Login)
		adminRoutes.POST("/change-password", controllers.ChangePassword)
		adminRoutes.POST("/reset-password", controllers.ResetPassword)
	}
	menuRoutes := r.Group("/menus")
	{
		menuRoutes.GET("/", controllers.MenuIndex)
	}
	vietqrRoutes := r.Group("/vietqr")
	{
		vietqrRoutes.GET("/create", controllers.CreateVietQRImage)
	}

	return r
}
