package router

import (
	// "project_alterra/controller"
	m "project_alterra/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo{
	e := echo.New()
	m.RemoveTrailingSlashMiddleware(e)
	e.Use(middleware.CORS())

	//Routing Movie
	movieGroup := e.Group("/api/v1")
	routesAll(movieGroup, e)
	
	m.LogMiddleware(e)

	return e
}