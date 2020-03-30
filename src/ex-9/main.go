package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello from the web side")
}

func main() {
	fmt.Println("Welcome to the server!")

	e := echo.New()

	e.GET("/", hello)

	e.Start(":8080")

}
