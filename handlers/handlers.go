package handlers

import (
	"net/http"
	"summeronsnow/web-echo/helpers"

	"github.com/labstack/echo/v4"
)

func HandleRequest(c echo.Context) error {
	var json interface{}
	var headers = make(map[string]string)
	var cookies = make(map[string]string)
	var params = make(map[string]string)

	c.Bind(&json)

	for key, values := range c.Request().Header {
		for _, value := range values {
			headers[key] = value
		}
	}

	for _, cookie := range c.Cookies() {
		cookies[cookie.Name] = cookie.Value
	}

	for key, values := range c.QueryParams() {
		for _, value := range values {
			params[key] = value
		}
	}

	r := &Response{
		Success: true,
		Data: Data{
			Method:      c.Request().Method,
			Path:        c.Path(),
			Headers:     headers,
			Cookies:     cookies,
			Payload:     json,
			QueryString: c.QueryString(),
			QueryParams: params,
			Host:        c.Request().Host,
			Ip:          c.RealIP(),
		},
		Metadata: Metadata{
			Timestamp: helpers.GetTimestamp(),
		},
	}

	return c.JSON(http.StatusOK, r)
}
