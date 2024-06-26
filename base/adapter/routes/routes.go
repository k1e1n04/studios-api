package routes

import (
	"github.com/k1e1n04/studios-api/base/adapter/middlewares"
	"github.com/k1e1n04/studios-api/src/adapter/api/auth"
	"github.com/k1e1n04/studios-api/src/adapter/api/study"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"os"
)

// InitRoutes は ルーティングを初期化
func InitRoutes(e *echo.Echo, container *dig.Container) {
	jwksURL := "https://cognito-idp." + os.Getenv("AWS_REGION") + ".amazonaws.com/" + os.Getenv("COGNITO_USER_POOL_ID") + "/.well-known/jwks.json"
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
	sg.Use(middlewares.CognitoJWTAuthMiddleware(jwksURL))
	sg.POST("/register", sc.Register)
	sg.GET("/list", sc.GetStudies)
	sg.GET("/:id", sc.GetStudy)
	sg.PUT("/update/:id", sc.UpdateStudy)
	sg.DELETE("/delete/:id", sc.DeleteStudy)
	sg.PUT("/review/complete/:id", sc.CompleteReview)
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
	tg.Use(middlewares.CognitoJWTAuthMiddleware(jwksURL))
	tg.GET("/list", tc.GetTags)

	// 認証コントローラ
	var ac auth.AuthController
	err = container.Invoke(func(c auth.AuthController) {
		ac = c
	})
	if err != nil {
		panic(err)
	}
	au := api.Group("/auth")
	au.POST("/signup", ac.SignUp)
	au.POST("/login", ac.Login)
}
