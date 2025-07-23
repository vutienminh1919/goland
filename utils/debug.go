package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Dump(c *gin.Context, data interface{}) {
	if !c.IsAborted() {
		c.IndentedJSON(http.StatusOK, data)
		c.Abort() // Dừng xử lý các middleware / handler sau
		panic("Dừng lại bởi utils.Dd")
	}
}
