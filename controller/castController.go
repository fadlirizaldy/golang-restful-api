package controller

import (
	"fmt"
	"net/http"

	"project_alterra/config"
	"project_alterra/model"

	"github.com/labstack/echo/v4"
)


func GetCastsController(c echo.Context) error {
	// deklarasi lagi tipe struct untuk tipe general
	DB := config.InitDB()
	type allCasts struct{
		Id       	int     `json:"id"`
		Name    	string  `json:"name"`
	}
	var casts []allCasts

	err := DB.Table("casts").Select("id","name").Find(&casts).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := Response{
		Success: true,
		Message: "Get all the casts",
		Data: casts,
	}
	return c.JSON(http.StatusOK, response)
}

func GetCastDetailController(c echo.Context) error {
	DB := config.InitDB()
	var casts []model.Cast

	id:= c.Param("id")

	err := DB.Where("id = ?", id).First(&casts).Error 
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := Response{
		Success: true,
		Message: "Get cast detail",
		Data: casts,
	}
	return c.JSON(http.StatusOK, response)
}

func CreateCastController(c echo.Context) error {
	DB := config.InitDB()
	casts := model.Cast{}
	c.Bind(&casts)

	err := DB.Save(&casts).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := Response{
		Success: true,
		Message: "Success add new cast",
		Data: casts,
	}
	return c.JSON(http.StatusOK, response)
}


// delete user
func DeleteCastByIdController(c echo.Context) error{
	DB := config.InitDB()
	casts := model.Cast{}
	// c.Bind(&casts)
	id:= c.Param("id")

	res := DB.Where("id = ?", id).Delete(&casts)
	// In sql, deleting non existed record not count as an error

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": res.Error,
		})
	} else if res.RowsAffected < 1 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": fmt.Sprintf("cast with id = %s doesn't exist", id),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success delete cast",
	})
}

//update user
func UpdateCastController(c echo.Context) error{
	DB := config.InitDB()
	casts := model.Cast{}
	c.Bind(&casts)

	id:= c.Param("id")

	err := DB.Where("id = ?", id).Updates(&casts).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	response := Response{
		Success: true,
		Message: "Success update user",
		Data: casts,
	}
	return c.JSON(http.StatusOK, response)
}