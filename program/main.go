package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.GET("/users", helloUser)
	e.Start(":8080")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func helloUser(c echo.Context) error {
	return c.String(http.StatusOK, "hello urnik")
}
