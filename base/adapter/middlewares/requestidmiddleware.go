package middlewares

import (
	"github.com/google/uuid"
	"github.com/k1e1n04/studios-api/base/config"
	"github.com/labstack/echo/v4"
)

// RequestIDMiddleware リクエストIDを付与するミドルウェア
func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqID := uuid.New().String()
		c.Request().Header.Set(echo.HeaderXRequestID, reqID)
		c.Response().Header().Set(echo.HeaderXRequestID, reqID)
		c.Set(config.RequestIdKey, reqID)
		return next(c)
	}
}
