package main

import (
	"gin-mvc/routes"
)

func main() {
	r := routes.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	} // Chạy tại localhost:8080
}
