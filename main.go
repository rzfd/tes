package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func EchoHandler(c echo.Context) error {
	fmt.Println("Hello World")
	return nil
}

func main() {
	e := echo.New()
	e.GET("/", EchoHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
