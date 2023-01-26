package controller

import (
	"fmt"
	"net/http"

	"project_alterra/config"
	"project_alterra/model"

	"github.com/labstack/echo/v4"
)

func GetMoviesController(c echo.Context) error {
	DB := config.InitDB()
	type allMovies struct{
		Id       	int     `json:"id"`
		Title    	string  `json:"title"`
	}
	
	var movies []allMovies

	err := DB.Table("movies").Select("id","title").Find(&movies).Error
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
	DB := config.InitDB()
	var movie []model.Movie

	id:= c.Param("id")

	err := DB.Preload("Casts").Where("id = ?", id).First(&movie).Error 

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
	DB := config.InitDB()
	movie := model.Movie{}
	c.Bind(&movie)

	err := DB.Save(&movie).Error

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
	DB := config.InitDB()
	movie := model.Movie{}
	// c.Bind(&movie)
	id:= c.Param("id")

	res := DB.Where("id = ?", id).Delete(&movie)
	// In sql, deleting non existed record not count as an error

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": res.Error,
		})
	} else if res.RowsAffected < 1 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": fmt.Sprintf("movie with id = %s doesn't exist", id),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete movie",
	})
}

//update user
func UpdateMovieController(c echo.Context) error{
	DB := config.InitDB()
	movie := model.Movie{}
	c.Bind(&movie)

	id:= c.Param("id")

	err := DB.Where("id = ?", id).Updates(&movie).Error
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