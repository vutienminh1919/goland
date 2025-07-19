package controllers

import (
	"gin-mvc/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeIndex(c *gin.Context) {
	message := services.GetWelcomeMessage()
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
