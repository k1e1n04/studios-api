package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/togisuma/standard-echo-serverless/src/adapter/api/example"
	"go.uber.org/dig"
)

func InitRoutes(e *echo.Echo, container *dig.Container) {
	// Exampleコントローラー
	var ec example.ExampleController
	err := container.Invoke(func(c example.ExampleController) {
		ec = c
	})
	if err != nil {
		panic(err)
	}

	// Example
	e.GET("/example/hello", ec.Hello)
}
