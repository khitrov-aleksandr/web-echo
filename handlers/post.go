package handlers

import (
	"net/http"
	"summeronsnow/web-echo/helpers"

	"github.com/labstack/echo/v4"
)

func Post(c echo.Context) error {
	var json interface{}

	c.Bind(&json)

	r := &Response{
		Success: true,
		Data: Data{
			YourData: json,
		},
		Metadata: Metadata{
			Timestamp: helpers.GetTimestamp(),
		},
	}

	return c.JSON(http.StatusOK, r)
}
