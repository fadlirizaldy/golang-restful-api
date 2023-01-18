package middleware

import (
	"github.com/labstack/echo/v4"
	appMiddleware "github.com/labstack/echo/v4/middleware"
)

func RemoveTrailingSlashMiddleware(e *echo.Echo) {
	e.Pre(appMiddleware.RemoveTrailingSlash())
}