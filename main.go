package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/togisuma/studios-api/base/adapter/api/errorhandler"
	"github.com/togisuma/studios-api/base/adapter/api/validator"
	"github.com/togisuma/studios-api/base/adapter/middlewares"
	"github.com/togisuma/studios-api/base/adapter/routes"
	"github.com/togisuma/studios-api/base/sharedkarnel/customlogger"
	"github.com/togisuma/studios-api/di"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var echoLambda *echoadapter.EchoLambda

// init Lambdaの初期化
func init() {
	e := initCommon()
	echoLambda = echoadapter.New(e)
}

// initLocalApp ローカル環境用の初期化
func initLocalApp() *echo.Echo {
	e := initCommon()
	return e
}

// main Lambdaのエントリーポイント
func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf(".env ファイルが存在しませんでした。")
	}
	env := os.Getenv("ENV")
	if env == "Local" {
		e := initLocalApp()
		e.Logger.Fatal(e.Start(":8080"))
	} else {
		lambda.Start(lambdaHandler)
	}
}

// lambdaHandler Lambdaのハンドラー
func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

// initCommon Lambdaとローカル環境の共通の初期化処理
func initCommon() *echo.Echo {
	// ロガーの作成
	customLogger, err := customlogger.NewCustomLogger()
	if err != nil {
		log.Fatalf("ロガー作成に失敗しました。: %v", err)
	}

	customLogger.Info("アプリケーションの初期化を開始します。")

	e := echo.New()

	setMiddleware(e, customLogger)
	setErrorHandler(e)
	// バリデーターの設定
	e.Validator = validator.NewCustomValidator()
	// エラーハンドラーの設定
	e.HTTPErrorHandler = errorhandler.HTTPErrorHandler

	// 依存関係の注入
	container := dig.New()
	err = di.RegisterDependencies(container, customLogger)
	if err != nil {
		customLogger.Panic("依存関係の注入に失敗しました。", zap.Error(err))
		panic(err)
	}

	// ルーティングの設定
	routes.InitRoutes(e, container)

	// 404ハンドラーの設定
	e.RouteNotFound("/*", func(c echo.Context) error { return c.NoContent(http.StatusNotFound) })

	return e
}

// setMiddleware ミドルウェアの設定
func setMiddleware(e *echo.Echo, logger *zap.Logger) {
	// リカバリーミドルウェア
	e.Use(middleware.Recover())
	// リクエストIDを付与するミドルウェア
	e.Use(middlewares.RequestIDMiddleware)
	// ロガーミドルウェア
	e.Use(middlewares.LoggingMiddleware(logger))
}

// エラーハンドラーの設定
func setErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		log.Printf("error: %v", err)
	}
}
