package router

import (
	"project_alterra/controller"

	"github.com/labstack/echo/v4"
)

func routesAll(group *echo.Group, e *echo.Echo){
	// Routing Movie
	group.GET("/movies", controller.GetMoviesController)
	group.GET("/movies/:id", controller.GetMovieDetailController)
	group.POST("/movies", controller.CreateMovieController)
	group.DELETE("/movies/:id", controller.DeleteMovieByIdController)
	group.PATCH("/movies/:id", controller.UpdateMovieController)
	
	//Routing Cast
	group.GET("/casts", controller.GetCastsController)
	group.GET("/casts/:id", controller.GetCastDetailController)
	group.POST("/casts", controller.CreateCastController)
	group.DELETE("/casts/:id", controller.DeleteCastByIdController)
	group.PATCH("/casts/:id", controller.UpdateCastController)
	
}