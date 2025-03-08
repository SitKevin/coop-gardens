package handlers

import (
	"net/http"

	"coop-gardens-be/usecase"

	"github.com/labstack/echo/v4"
)

// @Summary Login
// @Description Login
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Username and password"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /login [post]

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Login(c echo.Context) error {
	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	token, err := usecase.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
