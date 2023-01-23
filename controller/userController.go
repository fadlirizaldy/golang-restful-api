package controller

import (
	"net/http"
	"project_alterra/config"
	"project_alterra/helper"
	"project_alterra/middleware"
	"project_alterra/model"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

// 1. Buat register
// 2. Buat Login

func UserRegister(c echo.Context) error{
	DB := config.InitDB()
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	pwd, _ := helper.GeneratePassword(password)
	
	users := model.User{
		Name: name,
		Email: email,
		Password: pwd,
	}

	// check if the forms are fulfilled
	if name == "" || email == "" || password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "form not valid, please fill all the column",
		})
	}

	errEmail := DB.Where("email = ?", email).First(&users).Error
	if errEmail == nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Email already used",
		})
	}

	// Save the data if the email not used
	err := DB.Save(&users).Error

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
	pwd := c.FormValue("password")


	user := model.User{
		Email: email,
		Password: pwd,
	}

	DB := config.InitDB()
	err := DB.Where("email = ?", user.Email).First(&user).Error

	encryptionErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))

	if err != nil || encryptionErr != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to Login",
			"error": "Email or password is invalid",
		})
	}

	// JWT
	token, err := middleware.CreateToken(user.Id, user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to Login, Failed to generate Token",
			"error": err.Error(),
		})
	}

	userResponse := model.UserResponse{Id: user.Id, Name: user.Name, Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login!",
		"user" : userResponse,
	})
}