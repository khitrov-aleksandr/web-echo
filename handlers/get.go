package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	r := &Response{
		Success: true,
		Data:    "ok",
	}

	return c.JSON(http.StatusOK, r)
}
