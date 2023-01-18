package controller

import (
	"net/http"

	"project_alterra/config"
	"project_alterra/model"

	"github.com/labstack/echo/v4"
)

func GetMoviesController(c echo.Context) error {
	type allMovies struct{
		Id       	int     `json:"id"`
		Title    	string  `json:"title"`
	}
	
	var movies []allMovies

	err := config.DB.Table("movies").Select("id","title").Find(&movies).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := Response{
		Success: true,
		Message: "Get all movies",
		Data: movies,
	}
	return c.JSON(http.StatusOK, response)
}

func GetMovieDetailController(c echo.Context) error {
	var movie []model.Movie

	id:= c.Param("id")

	err := config.DB.Preload("Casts").Where("id = ?", id).First(&movie).Error 

	// err := config.DB.Where("id = ?", id).First(&movie).Error 
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := Response{
		Success: true,
		Message: "Get movie detail",
		Data: movie,
	}
	return c.JSON(http.StatusOK, response)
}

func CreateMovieController(c echo.Context) error {
	movie := model.Movie{}
	c.Bind(&movie)

	err := config.DB.Save(&movie).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	
	response := Response{
		Success: true,
		Message: "Get cast detail",
		Data: movie,
	}
	return c.JSON(http.StatusOK, response)
}


// delete user
func DeleteMovieByIdController(c echo.Context) error{
	movie := model.Movie{}
	c.Bind(&movie)
	id:= c.Param("id")

	err := config.DB.Where("id = ?", id).Delete(&movie).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success delete movie",
	})
}

//update user
func UpdateMovieController(c echo.Context) error{
	movie := model.Movie{}
	c.Bind(&movie)

	id:= c.Param("id")

	err := config.DB.Where("id = ?", id).Updates(&movie).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	
	response := Response{
		Success: true,
		Message: "Success update movie",
		Data: movie,
	}
	return c.JSON(http.StatusOK, response)
}