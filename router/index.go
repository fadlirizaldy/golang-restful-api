package router

import (
	"project_alterra/constants"
	"project_alterra/controller"
	m "project_alterra/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo{
	e := echo.New()
	m.RemoveTrailingSlashMiddleware(e)
	e.Use(middleware.CORS())

	//Routing User login and Register
	e.PUT("/register", controller.UserRegister)
	e.POST("/login", controller.UserLogin)

	//Routing all
	movieGroup := e.Group("/api/v1")
	movieGroup.Use(echojwt.JWT([]byte(constants.SECRET_JWT)))
	routesAll(movieGroup)
	
	m.LogMiddleware(e)

	return e
}