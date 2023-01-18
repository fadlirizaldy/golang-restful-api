package controller

import (
	"net/http"

	"project_alterra/config"
	"project_alterra/model"

	"github.com/labstack/echo/v4"
)


func GetCastsController(c echo.Context) error {
	// deklarasi lagi tipe struct untuk tipe general
	type allCasts struct{
		Id       	int     `json:"id"`
		Name    	string  `json:"name"`
	}
	var casts []allCasts

	err := config.DB.Table("casts").Select("id","name").Find(&casts).Error
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
	var casts []model.Cast

	id:= c.Param("id")

	err := config.DB.Where("id = ?", id).First(&casts).Error 
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
	casts := model.Cast{}
	c.Bind(&casts)

	err := config.DB.Save(&casts).Error

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
	casts := model.Cast{}
	c.Bind(&casts)
	id:= c.Param("id")

	err := config.DB.Where("id = ?", id).Delete(&casts).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success delete cast",
	})
}

//update user
func UpdateCastController(c echo.Context) error{
	casts := model.Cast{}
	c.Bind(&casts)

	id:= c.Param("id")

	err := config.DB.Where("id = ?", id).Updates(&casts).Error
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