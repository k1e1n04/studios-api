package routes

import (
	"github.com/k1e1n04/studios-api/src/adapter/api/example"
	"github.com/labstack/echo/v4"
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
