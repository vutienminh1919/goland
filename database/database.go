package database

import (
	"fmt"
	"gin-mvc/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnv("DB_USER", "root"),
		config.GetEnv("DB_PASS", ""),
		config.GetEnv("DB_HOST", "127.0.0.1"),
		config.GetEnv("DB_PORT", "3306"),
		config.GetEnv("DB_NAME", "laravel"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Không thể kết nối database: " + err.Error())
	}

	fmt.Println("✅ Đã kết nối thành công MySQL")
	DB = database
}
