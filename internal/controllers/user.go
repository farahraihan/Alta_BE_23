package controllers

import (
	"net/http"

	"taskday7.4/configs"
	"taskday7.4/internal/models"

	"github.com/labstack/echo/v4"
)

// Handler untuk registrasi user
func RegisterUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	configs.DB.Create(&user)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User successfully registered",
	})
}

// Handler untuk login user
func LoginUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := configs.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid username or password",
		})
	}
	// Generate JWT token
	// Implementasi token JWT bisa ditambahkan di sini
	return c.JSON(http.StatusOK, map[string]string{
		"token": "jwt_token_here",
	})
}
