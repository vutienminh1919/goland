package controllers

import (
	"gin-mvc/database"
	"gin-mvc/models"
	"gin-mvc/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

func HomeIndex(c *gin.Context) {
	message := services.GetWelcomeMessage()
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

type SearchRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAllUsersByParams(c *gin.Context) {

	var input SearchRequest
	var users []models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ", "detail": err.Error()})
		return
	}

	query := database.DB

	if input.Name != "" {
		query = query.Where("name LIKE ?", "%"+input.Name+"%")
	}
	if input.Email != "" {
		query = query.Where("email LIKE ?", "%"+input.Email+"%")
	}

	result := query.Find(&users)

	var response []map[string]interface{}
	for _, user := range users {
		role := "member"
		if strings.EqualFold(user.Name, "admin") {
			role = "admin"
		}
		userMap := map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  role, // thêm trường bất kỳ
		}

		response = append(response, userMap)

	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}
	if input.Password == "" {
		input.Password = "123456"
	}

	// Hash mật khẩu
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể mã hóa mật khẩu"})
		return
	}

	// Tạo user mới
	user := models.User{
		Name:        input.Name,
		Email:       input.Email,
		Password:    string(hashedPassword),
		TimeCreated: int(time.Now().Unix()),
		TimeUpdated: int(time.Now().Unix()),
	}

	// Lưu vào DB
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo người dùng", "detail": err.Error()})
		return
	}

	// Trả kết quả
	c.JSON(http.StatusCreated, gin.H{
		"message": "Tạo người dùng thành công",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})

}
