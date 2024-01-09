package middlewares

import (
	"github.com/k1e1n04/studios-api/base/config"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/customlogger"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// LoggingMiddleware ロガーミドルウェア
func LoggingMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// 元のロガーインスタンスを保存
			originalLogger := logger

			// リクエストIDがあればロガーを新たに生成
			reqID, exists := c.Get(config.RequestIdKey).(string)
			if exists {
				logger = customlogger.WithRequestID(logger, reqID)
				defer func() {
					// リクエストが終了したら元のロガーに戻す
					logger = originalLogger
				}()
			}

			// 新しいロガーインスタンスをコンテキストにセット
			c.Set(config.LoggerKey, logger)

			// 次のミドルウェアやルートハンドラを実行
			err := next(c)

			// ログ情報を収集
			req := c.Request()
			res := c.Response()
			latency := time.Since(start)
			ip := c.RealIP()

			logger.Info("Request details",
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.String("ip", ip),
				zap.Int("status", res.Status),
				zap.Duration("latency", latency),
			)

			return err
		}
	}
}
