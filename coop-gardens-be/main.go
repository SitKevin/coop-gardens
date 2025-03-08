package main

import (
	"log"

	"coop-gardens-be/config"
	_ "coop-gardens-be/docs"
	"coop-gardens-be/internal/api/routers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Kết nối DB
	config.InitDB()
	// Khởi tạo routes
	routers.LoginRoutes(e)

	log.Println("🚀 Server đang chạy tại: http://localhost:8080")
	e.Start(":8080")
}
