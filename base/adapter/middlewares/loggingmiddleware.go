package middlewares

import (
	"github.com/togisuma/standard-echo-serverless/base/config"
	"github.com/togisuma/standard-echo-serverless/base/sharedkarnel/customlogger"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// LoggingMiddleware ロガーミドルウェア
func LoggingMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			// カスタムロガーでリクエストID出力するようにし、コンテキストに追加
			reqID, exists := c.Get(config.RequestIdKey).(string)
			if exists {
				logger = customlogger.WithRequestID(logger, reqID)
			}
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
