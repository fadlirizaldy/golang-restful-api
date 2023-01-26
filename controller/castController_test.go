package controller

import (
	"bytes"
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"
	"net/http/httptest"
	"project_alterra/config"
	"strconv"
	"testing"
	"time"

	// "project_alterra/helper"
	"project_alterra/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)



type ResponseData struct {
	Success	bool `json:"success"`
	Message	string `json:"message"`
	Data	model.Cast `json:"data"`
}

func CreateCast(name string, birth_place string, birthday time.Time, rating int) *model.Cast{
	DB := config.InitDB()
	cast := model.Cast{
		Name: name,
		Birth_place: birth_place,
		Birthday: birthday,
		Rating: rating,
	}

	if err := DB.Table("casts").Create(&cast).Error; err != nil {
		return nil
	}

	return &cast
}

func TestGetCastsController (t *testing.T){
	config.InitMigrate()

	casts := []model.Cast{
		{
			Name: "Kakashi", Birth_place: "Tokyo", Birthday: time.Now(), Rating: 5,
		},
		{
			Name: "Kushina", Birth_place: "Nagasaki", Birthday: time.Now(), Rating: 4,
		},
	}

	for _, cast := range casts {
		CreateCast(cast.Name, cast.Birth_place, cast.Birthday, cast.Rating)
	}

	testCases := []struct {
		name       string
		path       string
		expectCode int
		dataSize   int
	}{
		{
			name:       "get all casts",
			path:       "/api/v1/casts",
			expectCode: http.StatusOK,
			dataSize:   len(casts),
		},
	}

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/casts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetCastsController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response struct {
				Casts []model.Cast `json:"data"`
			}

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, len(response.Casts)>0, testCase.dataSize > 0)
		}
	}
}

func TestGetCastDetailController(t *testing.T){
	config.InitMigrate()

	cast := CreateCast("Sasuke", "Konoha", time.Now(), 5)

	testCases := []struct {
		name       string
		path       string
		castId    string
		expectCode int
	}{
		{
			name:       "get detail cast",
			path:       "/api/v1/casts",
			castId:    strconv.Itoa(int(cast.Id)),
			expectCode: http.StatusOK,
		},
	}

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/casts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.castId)
	

	if assert.NoError(t, GetCastDetailController(c)) {
		assert.Equal(t, testCase.expectCode, rec.Code)
		body := rec.Body.String()

		var response ResponseData

		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, true, response.Success)
		}
	}
}

func TestCreateCastController(t *testing.T){
	config.InitMigrate()

	testCases := []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "add new cast",
			path:       "/api/v1/casts",
			expectCode: http.StatusOK,
		},
	}

	reqBody := map[string]interface{}{
		"name":  "avatar kora",
		"birth_place":   "Sudan",
		"birthday": time.Now(),
		"rating": 5,
	}
	reqBodyJSON, _ := json.Marshal(reqBody)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/casts", bytes.NewReader(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateCastController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response ResponseData

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, true, response.Success)
		}
	}

}

func TestUpdateCastController(t *testing.T){
	config.InitMigrate()

	cast := CreateCast("Sasuke", "Konoha", time.Now(), 5)

	testCases := []struct {
		name       string
		path       string
		castId    string
		expectCode int
	}{
		{
			name:       "Update cast",
			path:       "/api/v1/casts",
			castId:    strconv.Itoa(int(cast.Id)),
			expectCode: http.StatusOK,
		},
	}

	reqBody := map[string]interface{}{
		"name":  "Sasuke Updated",
		"birth_place":"Pindahan Sudan",
		"birthday": time.Now(),
		"rating": 3,
	}
	reqBodyJSON, _ := json.Marshal(reqBody)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/casts", bytes.NewReader(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.castId)
	

	if assert.NoError(t, UpdateCastController(c)) {
		assert.Equal(t, testCase.expectCode, rec.Code)
		body := rec.Body.String()

		var response ResponseData

		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, true, response.Success)
		}
	}
}

func TestDeleteCastController(t *testing.T){
	config.InitMigrate()

	cast := CreateCast("Sasuke", "Konoha", time.Now(), 5)

	testCases := []struct {
		name       string
		path       string
		castId    string
		expectCode int
	}{
		{
			name:       "delete cast data",
			path:       "/api/v1/casts",
			castId:    strconv.Itoa(int(cast.Id)),
			expectCode: http.StatusOK,
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/casts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.castId)
	

	if assert.NoError(t, DeleteCastByIdController(c)) {
		assert.Equal(t, testCase.expectCode, rec.Code)
		body := rec.Body.String()
		fmt.Println(body)

		var response struct{
			Message string `json:"message"`
		}

		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, "Success delete cast", response.Message)
		}
	}
}

