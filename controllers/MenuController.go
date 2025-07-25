package controllers

import (
	"encoding/json"
	"gin-mvc/database"
	"gin-mvc/models"
	"gin-mvc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MenuIndex(c *gin.Context) {
	var menus []models.Menu
	query := database.DB

	err := query.Find(&menus).Error
	if err != nil {
		utils.LogError("Lỗi khi lấy danh sách menu: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy danh sách menu"})
		return
	}

	// Ghi log menus dạng JSON đẹp
	menusJSON, err := json.MarshalIndent(menus, "", "  ")
	if err != nil {
		utils.LogError("Không thể format dữ liệu menus để log")
	} else {
		utils.LogInfo("Danh sách menus trả về:\n" + string(menusJSON))
	}
	c.JSON(http.StatusOK, gin.H{
		"menus": menus,
	})

}
