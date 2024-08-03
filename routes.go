package main

import (
	"summeronsnow/web-echo/handlers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.Any("/", handlers.HandleRequest)
}
