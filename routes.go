package main

import (
	"summeronsnow/web-echo/handlers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/", handlers.Get)
	e.POST("/", handlers.Post)
}
