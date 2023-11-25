package routes

import (
	"github.com/k1e1n04/studios-api/src/adapter/api/study"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

func InitRoutes(e *echo.Echo, container *dig.Container) {
	api := e.Group("/api/v1")

	// 学習コントローラ
	var sc study.StudyController
	err := container.Invoke(func(c study.StudyController) {
		sc = c
	})
	if err != nil {
		panic(err)
	}

	sg := api.Group("/study")
	sg.POST("/register", sc.Register)
	sg.GET("/list", sc.GetStudies)
}
