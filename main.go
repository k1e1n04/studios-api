package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/k1e1n04/studios-api/base/adapter/api/errorhandler"
	"github.com/k1e1n04/studios-api/base/adapter/api/validator"
	"github.com/k1e1n04/studios-api/base/adapter/middlewares"
	"github.com/k1e1n04/studios-api/base/adapter/routes"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/customlogger"
	"github.com/k1e1n04/studios-api/di"
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

// Secret は RDSのユーザー名とパスワードを格納する構造体
type Secret struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// init Lambdaの初期化
func init() {
	e := initCommon()
	echoLambda = echoadapter.New(e)
}

// main Lambdaのエントリーポイント
func main() {
	lambda.Start(lambdaHandler)
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

	e := echo.New()

	setMiddleware(e, customLogger)
	setErrorHandler(e)
	// バリデーターの設定
	e.Validator = validator.NewCustomValidator()
	// エラーハンドラーの設定
	e.HTTPErrorHandler = errorhandler.HTTPErrorHandler

	// TODO: https://docs.aws.amazon.com/ja_jp/secretsmanager/latest/userguide/retrieving-secrets_lambda.html に移行する
	//secret, err := getSecrets(os.Getenv("DB_SECRET_NAME"))
	//if err != nil {
	//	log.Fatalf("Failed to get secret: %v", err)
	//}

	// 依存関係の注入
	container := dig.New()
	err = di.RegisterDependencies(container, customLogger, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	if err != nil {
		customLogger.Panic("依存関係の注入に失敗しました。", zap.Error(err))
		panic(err)
	}

	// CORSの設定 TODO: 本番環境では許可するオリジンを絞る
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

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

// getSecrets SecretsManagerからシークレットを取得
func getSecrets(secretName string) (*Secret, error) {
	// SecretsManager のセッションを作成
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	// SecretsManager のクライアントを作成
	svc := secretsmanager.New(sess)

	// シークレットを取得
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return nil, err
	}

	// SecretStringをパース
	var secret Secret
	err = json.Unmarshal([]byte(*result.SecretString), &secret)
	if err != nil {
		return nil, err
	}

	return &secret, nil
}
