package middlewares

import (
	"github.com/k1e1n04/studios-api/base"
	"github.com/k1e1n04/studios-api/base/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strings"
)

// APIKeyAuthenticationMiddleware は APIキー認証のミドルウェア
func APIKeyAuthenticationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// ロガーを取得
			logger := c.Get(config.LoggerKey).(*zap.Logger)
			// Authorizationヘッダーからトークンを取得
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				logger.Warn("Authorizationヘッダーは必須です。")
				return echo.NewHTTPError(http.StatusUnauthorized, base.AuthenticationHeaderRequired)
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				logger.Warn("Authorizationヘッダーの形式が正しくありません。")
				return echo.NewHTTPError(http.StatusUnauthorized, base.InvalidAuthenticationHeader)
			}

			apiKey := parts[1]

			// APIキーを検証
			if apiKey != os.Getenv("API_KEY") {
				logger.Warn("APIキーの検証に失敗しました。")
				return echo.NewHTTPError(http.StatusUnauthorized, base.InvalidAPIKey)
			}

			return next(c)
		}
	}
}
