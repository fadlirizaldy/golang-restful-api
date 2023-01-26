package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"project_alterra/config"
	"project_alterra/helper"
	"project_alterra/model"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)


type ResponseSuccessRegister struct {
	Success 	bool 	`json:"success"`
	Message		string	`json:"message"`
}

type ResponseSuccessLogin struct {
	Message		string	`json:"message"`
	UserData	model.UserResponse `json:"data"`
}

func CreateUser(name, email, password string) *model.User {
	DB := config.InitDB()
	pwd, _ := helper.GeneratePassword(password)
	user := model.User{
		Name:     name,
		Email:    email,
		Password: pwd,
	}

	if err := DB.Save(&user).Error; err != nil {
		return nil
	}

	return &user
}

func TruncateUsersTable() {
	DB := config.InitDB()
	// DB.Exec("DELETE FROM users ORDER BY id DESC LIMIT 1")
	DB.Exec("TRUNCATE TABLE users")
}

func GetToken() string {
	user := CreateUser("nanang", "naning@ymail.go.hk", "test123")

	data := url.Values{}
    data.Set("email", user.Email)
	data.Set("password", user.Password)

	fmt.Println("======DATA=====")
	fmt.Println(data)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/login")
	UserLogin(c)

	body := rec.Body.String()

	var responseLogin struct {
		User model.UserResponse `json:"data"`
	}

	_ = json.Unmarshal([]byte(body), &responseLogin)

	fmt.Println("=========Body=======")
	fmt.Println(body)
	fmt.Println("=====================")
	fmt.Println(responseLogin.User.Token)

	return fmt.Sprintf("Bearer %s", responseLogin.User.Token)
}

// Test User Register
func TestUserRegister(t *testing.T) {
	config.InitMigrate()

	testCases := []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "create normal user",
			path:       "/register",
			expectCode: http.StatusOK,
		},
	}

	e := echo.New()

	data := url.Values{}
    data.Set("name", "kakanda")
    data.Set("email", "kakanda@ymail.com")
	data.Set("password", "kakanda123")
	
	req := httptest.NewRequest(http.MethodPut, "/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, UserRegister(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response ResponseSuccessRegister

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, response.Message, "Success register new account")
		}
	}

	TruncateUsersTable()
}

func TestLoginUser (t *testing.T){
	config.InitMigrate()
	CreateUser("nanang", "naning@ymail.go.hk", "test123")
	testCases2 := []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "Normal login user",
			path:       "/login",
			expectCode: http.StatusOK,
		},
	}

	e := echo.New()

	dataLogin := url.Values{}
    dataLogin.Set("email", "naning@ymail.go.hk")
	dataLogin.Set("password", "test123")

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(dataLogin.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases2 {
		c.SetPath(testCase.path)

		if assert.NoError(t, UserLogin(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response ResponseSuccessLogin

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, response.Message, "Success Login!")
		}
	}
	TruncateUsersTable()
}