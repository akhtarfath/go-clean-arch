package main

import (
	_appMiddleware "github.com/bpti-uhamka/uhamka-api/app/middleware"
	"github.com/labstack/echo"
)

func middleware() *echo.Echo {
	// Echo instance
	e := echo.New()

	// middleware for database
	middleware := _appMiddleware.InitMiddleware()
	e.Use(middleware.CORS) // CORS middleware

	return e
}
