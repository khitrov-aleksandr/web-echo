package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Response struct {
	Success  bool     `json:"success"`
	Data     any      `json:"data"`
	Metadata Metadata `json:"metadata"`
}

type Data struct {
	YourData any `json:"yourData"`
}

type Metadata struct {
	Timestamp int `json:"timestamp"`
}

var json interface{}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", getHandler)
	e.POST("/", postHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func getHandler(c echo.Context) error {
	r := &Response{
		Success: true,
		Data:    "ok",
	}

	return c.JSON(http.StatusOK, r)
}

func postHandler(c echo.Context) error {
	c.Bind(&json)

	r := &Response{
		Success: true,
		Data: Data{
			YourData: json,
		},
		Metadata: Metadata{
			Timestamp: 10,
		},
	}

	return c.JSON(http.StatusOK, r)
}
