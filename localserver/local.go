package main

import (
	"github.com/joho/godotenv"
	"github.com/k1e1n04/gosmm/v2/pkg/gosmm"
	"github.com/k1e1n04/studios-api/base/adapter/api/errorhandler"
	"github.com/k1e1n04/studios-api/base/adapter/api/validator"
	"github.com/k1e1n04/studios-api/base/adapter/middlewares"
	"github.com/k1e1n04/studios-api/base/adapter/routes"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/customlogger"
	"github.com/k1e1n04/studios-api/di"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
)

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
	migrate()
	e := initLocalApp()
	e.Logger.Fatal(e.Start(":8080"))
}

// initCommon Lambdaとローカル環境の共通の初期化処理
func initCommon() *echo.Echo {
	// ロガーの作成
	customLogger, err := customlogger.NewCustomLogger()
	if err != nil {
		log.Fatalf("ロガー作成に失敗しました。: %v", err)
	}

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

// migrate マイグレーションを実行
func migrate() {
	driver := "mysql"
	config := gosmm.DBConfig{
		Driver:   driver,
		Host:     os.Getenv("DB_HOST"),
		Port:     3306,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := gosmm.ConnectDB(config)
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	err = gosmm.Migrate(db, os.Getenv("MIGRATIONS_DIR"), driver)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	err = gosmm.CloseDB(db)
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
}
