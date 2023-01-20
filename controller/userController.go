package controller

import (
	// "fmt"
	"net/http"
	"project_alterra/config"
	"project_alterra/middleware"
	"project_alterra/model"

	"github.com/labstack/echo/v4"
)

// 1. Buat register
// 2. Buat Login

func UserRegister(c echo.Context) error{
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	users := model.User{
		Name: name,
		Email: email,
		Password: password,
	}

	if name == "" || email == "" || password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "form not valid, please fill all the column",
		})
	}

	err := config.DB.Save(&users).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	response := echo.Map{
		"success": true,
		"message": "Success register new account",
	}
	return c.JSON(http.StatusOK, response)

}

func UserLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := model.User{
		Email: email,
		Password: password,
	}

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	// fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to Login",
			"error": err.Error(),
		})
	}

	// JWT
	token, err := middleware.CreateToken(user.Id, user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to Login",
			"error": err.Error(),
		})
	}

	userResponse := model.UserResponse{Id: user.Id, Name: user.Name, Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success Login!",
		"user" : userResponse,
	})
}