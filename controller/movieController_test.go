package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project_alterra/config"
	"project_alterra/model"
	"strconv"
	"testing"

	// "project_alterra/helper"
	// "project_alterra/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func CreateMovie(title, language, status string, rating float32) *model.Movie {
	DB := config.InitDB()
	movie := model.Movie{
		Title: title,
		Language: language,
		Status: status,
		Rating: rating,
	}

	if err := DB.Table("movies").Create(&movie).Error; err != nil {
		return nil
	}

	return &movie
}

func TestGetMoviesController(t *testing.T){
	config.InitMigrate()

	testCases := []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get all data movies",
			path:       "/api/v1/movies",
			expectCode: http.StatusOK,
		},
	}

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/movies", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetMoviesController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response ResponseData

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Get all movies", response.Message)
		}
	}
}

func TestGetMovieDetailController(t *testing.T){
	config.InitMigrate()

	movie := CreateMovie("Ashiap Man", "Sundanese", "Ended", 4.3)

	testCases := []struct {
		name       string
		path       string
		castId    string
		expectCode int
	}{
		{
			name:       "get detail movie",
			path:       "/api/v1/movies",
			castId:    strconv.Itoa(int(movie.Id)),
			expectCode: http.StatusOK,
		},
	}

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/movies", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.castId)
	

	if assert.NoError(t, GetMovieDetailController(c)) {
		assert.Equal(t, testCase.expectCode, rec.Code)
		body := rec.Body.String()
		fmt.Println(body)

		var response ResponseData

		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			assert.Error(t, err, "error")
		}

		fmt.Println(response)
		assert.Equal(t, true, response.Success)
		}
	}	
}

func TestCreateMovieController(t *testing.T){
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
		"title": "Avatar Aang Tuah",
		"language": "Arab",
		"status": "Ongoing",
		"rating": 4.8,
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

		if assert.NoError(t, CreateMovieController(c)) {
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

func TestUpdateMovieController(t *testing.T){
	config.InitMigrate()

	movie := CreateMovie("Ashiap Man II", "Batak", "Ongoing", 4.3)

	testCases := []struct {
		name       string
		path       string
		castId    string
		expectCode int
	}{
		{
			name:       "update movie",
			path:       "/api/v1/movies",
			castId:    strconv.Itoa(int(movie.Id)),
			expectCode: http.StatusOK,
		},
	}

	reqBody := map[string]interface{}{
		"title": "Ashiap Man II Updated",
		"language": "Batavianese",
		"status": "ended Updated",
		"rating": 5.0,
	}
	reqBodyJSON, _ := json.Marshal(reqBody)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/movies", bytes.NewReader(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.castId)
	

	if assert.NoError(t, UpdateMovieController(c)) {
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

func TestDeleteMovieController(t *testing.T){
	config.InitMigrate()

	movie := CreateMovie("Ashiap Man II", "Batak", "Ongoing", 4.3)

	testCases := []struct {
		name       string
		path       string
		movieId    string
		expectCode int
	}{
		{
			name:       "delete movie data",
			path:       "/api/v1/movies",
			movieId:    strconv.Itoa(int(movie.Id)),
			expectCode: http.StatusOK,
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/movies", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.movieId)
	

	if assert.NoError(t, DeleteMovieByIdController(c)) {
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

		assert.Equal(t, "Success delete movie", response.Message)
		}
	}
}