package main

import (
	"fmt"
	"gin-mvc/config"
	"gin-mvc/database"
	"gin-mvc/redis"
	"gin-mvc/routes"
	"gin-mvc/scheduler"
	"gin-mvc/utils"
)

func main() {

	database.Connect() // <--- kết nối DB trước
	redis.ConnectRedis()
	config.LoadEnv()
	utils.InitLogger()

	port := config.GetEnv("APP_PORT", "8080")
	fmt.Println("App đang chạy tại cổng:", port)
	scheduler.StartScheduler()

	r := routes.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
