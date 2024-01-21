package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/k1e1n04/studios-api/base"
	"github.com/k1e1n04/studios-api/base/config"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// CognitoJWTAuthMiddleware は CognitoのJWT認証のミドルウェア
func CognitoJWTAuthMiddleware(jwksURL string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// ロガーを取得
			logger := c.Get(config.LoggerKey).(*zap.Logger)
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				logger.Warn("Authorizationは必須です。")
				return echo.NewHTTPError(http.StatusUnauthorized, base.AuthenticationHeaderRequired)
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				logger.Warn("Authorizationヘッダーの形式が正しくありません。")
				return echo.NewHTTPError(http.StatusUnauthorized, base.InvalidAuthenticationHeader)
			}

			accessToken := parts[1]

			token, err := verifyCognitoToken(jwksURL, accessToken, logger)
			if err != nil {
				logger.Warn(err.Error())
				return echo.NewHTTPError(http.StatusUnauthorized, base.InvalidToken)
			}

			// User IDをContextにセット
			userID, err := getUserIDByToken(token, logger)
			if err != nil {
				logger.Warn(err.Error())
				return echo.NewHTTPError(http.StatusUnauthorized, base.InvalidToken)
			}
			c.Set(config.UserIDKey, userID)
			return next(c)
		}
	}
}

// verifyCognitoToken は Cognitoのトークンを検証する関数
func verifyCognitoToken(jwksURL string, accessToken string, logger *zap.Logger) (jwt.Token, error) {
	// JWTの検証に使用するJWKセットを取得
	jwkSet, err := jwk.Fetch(context.Background(), jwksURL)
	if err != nil {
		logger.Warn("JWKセットの取得に失敗しました")
		return nil, fmt.Errorf("JWKセットの取得に失敗しました")
	}

	// トークンの検証
	token, err := jwt.ParseString(accessToken, jwt.WithKeySet(jwkSet), jwt.WithValidate(true))
	if err != nil {
		logger.Warn("トークンの解析と検証に失敗しました")
		return nil, fmt.Errorf("トークンの解析と検証に失敗しました")
	}

	return token, nil
}

// トークンからユーザー情報を取得する
func getUserIDByToken(token jwt.Token, logger *zap.Logger) (string, error) {
	claims, err := token.AsMap(context.TODO())
	if err != nil {
		logger.Warn("トークンからクレームを取得できませんでした。", zap.Error(err))
		return "", err
	}

	userIDValue, ok := claims["sub"]
	if !ok || userIDValue == nil {
		err = errors.New("クレーム 'sub' が見つからないか、nil です。")
		logger.Warn(err.Error(), zap.Error(err))
		return "", err
	}

	userID, ok := userIDValue.(string)
	if !ok {
		err = errors.New("'sub' クレームがstring型ではありません。")
		logger.Warn(err.Error(), zap.Error(err))
		return "", err
	}
	println(userID)

	return userID, nil
}
