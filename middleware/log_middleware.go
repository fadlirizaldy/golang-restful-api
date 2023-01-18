package middleware

import (
	"github.com/labstack/echo/v4"
	appMiddleware "github.com/labstack/echo/v4/middleware"
)

func LogMiddleware(e *echo.Echo) {
	e.Use(appMiddleware.LoggerWithConfig(appMiddleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
	}))
}