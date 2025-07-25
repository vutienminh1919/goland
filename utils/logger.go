package utils

import (
	"log"
	"os"
)

var logger *log.Logger

// InitLogger khởi tạo logger (chỉ gọi 1 lần, ví dụ trong main.go)
func InitLogger() {
	logFilePath := "D:/GOLANG/DEV TEST/gin-mvc/log/log.txt"

	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Không thể mở file log: %v", err)
	}

	// Tạo logger riêng để dùng global
	logger = log.New(f, "", log.LstdFlags|log.Lshortfile)
}

// LogInfo ghi log thông tin
func LogInfo(message string) {
	if logger != nil {
		logger.Println("[INFO]", message)
	}
}

// LogError ghi log lỗi
func LogError(message string) {
	if logger != nil {
		logger.Println("[ERROR]", message)
	}
}
