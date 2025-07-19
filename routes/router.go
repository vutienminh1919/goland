package routes

import (
	"gin-mvc/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/", controllers.HomeIndex)

	return r
}
