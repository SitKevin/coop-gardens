package main

import (
	"log"

	"coop-gardens-be/config"
	"coop-gardens-be/internal/api/routers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load biến môi trường từ .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	// Kết nối DB
	config.InitDB()
	// Khởi tạo routes
	routers.LoginRoutes(e)

	e.Start(":8080")
}
