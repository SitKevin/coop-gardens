package routers

import (
	"coop-gardens-be/internal/api/handlers"

	"github.com/labstack/echo/v4"
)

func LoginRoutes(e *echo.Echo) {
	e.POST("/login", handlers.Login)
}
