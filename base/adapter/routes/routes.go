package routes

import (
	"github.com/k1e1n04/studios-api/base/adapter/middlewares"
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
	sg.Use(middlewares.APIKeyAuthenticationMiddleware())
	sg.POST("/register", sc.Register)
	sg.GET("/list", sc.GetStudies)
	sg.GET("/:id", sc.GetStudy)
	sg.PUT("/update/:id", sc.UpdateStudy)
	sg.DELETE("/delete/:id", sc.DeleteStudy)
	sg.POST("/review/complete/:id", sc.CompleteReview)
	sg.GET("/review/list", sc.GetStudiesReview)

	// タグコントローラ
	var tc study.TagController
	err = container.Invoke(func(c study.TagController) {
		tc = c
	})
	if err != nil {
		panic(err)
	}
	tg := api.Group("/tag")
	tg.Use(middlewares.APIKeyAuthenticationMiddleware())
	tg.GET("/list", tc.GetTags)
}
