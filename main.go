package main

import (
	"fmt"
	"gin-mvc/config"
	"gin-mvc/database"
	"gin-mvc/routes"
)

func main() {

	database.Connect() // <--- kết nối DB trước
	config.LoadEnv()

	port := config.GetEnv("APP_PORT", "8080")
	fmt.Println("App đang chạy tại cổng:", port)

	r := routes.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}

	//r := routes.SetupRouter()
	//err := r.Run(":8080")
	//if err != nil {
	//	return
	//} // Chạy tại localhost:8080
}
