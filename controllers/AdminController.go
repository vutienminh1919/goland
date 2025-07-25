package controllers

import (
	"fmt"
	"gin-mvc/database"
	"gin-mvc/models"
	"gin-mvc/redis"
	"gin-mvc/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var input LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email không tồn tại"})
		return
	}

	// So sánh password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Mật khẩu không đúng"})
		return
	}

	// Tạo token JWT
	token, err := utils.GenerateJWT(uint(user.ID), user.Email, user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không tạo được token"})
		return
	}

	redisKey := fmt.Sprintf("token:%d", user.ID)
	err = redis.Rdb.Set(redis.Ctx, redisKey, token, time.Hour*1).Err() // TTL = 1 giờ
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không lưu được token vào Redis"})
		return
	}
	log.Println("🔐 Token đã lưu Redis:", redisKey, token)

	utils.LogInfo(fmt.Sprintf("Login success - Email: %s, UserID: %d", user.Email, user.ID))
	c.JSON(http.StatusOK, gin.H{
		"message": "Đăng nhập thành công",
		"token":   token,
	})
}

type ResetRequest struct {
	Email string `json:"email"`
}

func ResetPassword(c *gin.Context) {
	var input ResetRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email không tồn tại"})
		return
	}
	newPassword := "123456"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể má hóa mật khẩu"})
		return
	}
	user.Password = string(hashedPassword)
	user.TimeUpdated = int(time.Now().Unix())
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể cập nhật mật khẩu"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Mật khẩu đã được đặt lại về 123456",
	})

}

type ChangePasswordRequest struct {
	Email           string `json:"email"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func ChangePassword(c *gin.Context) {
	var input ChangePasswordRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email không tồn tại"})
		return
	}
	if input.NewPassword != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Xác nhân mật khẩu không trùng khớp"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể mã hóa mật khẩu mới"})
		return
	}

	user.Password = string(hashedPassword)
	user.TimeUpdated = int(time.Now().Unix())

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể cập nhật mật khẩu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Đổi mật khẩu thành công"})

}
