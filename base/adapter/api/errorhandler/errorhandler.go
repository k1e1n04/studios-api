package errorhandler

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/k1e1n04/studios-api/base"
	"github.com/k1e1n04/studios-api/base/config"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

// HTTPErrorHandler は HTTPエラーハンドラー
func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := base.InternalServerError
	// エラーログを出力
	logger := c.Get(config.LoggerKey).(*zap.Logger)

	var e interface{}

	// HTTPErrorのチェック
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
	} else if errors.As(err, &e) {
		switch e := e.(type) {
		case *customerrors.BadRequestError:
			logger.Warn(e.DebugMessage, zap.Error(e.Cause))
			code = http.StatusBadRequest
			msg = e.Error()
		case *customerrors.NotFoundError:
			logger.Warn(e.DebugMessage, zap.Error(e.Cause))
			code = http.StatusNotFound
			msg = e.Error()
		case *customerrors.ConflictError:
			logger.Warn(e.DebugMessage, zap.Error(e.Cause))
			code = http.StatusConflict
			msg = e.Error()
		case *customerrors.UnauthorizedError:
			logger.Warn(e.DebugMessage, zap.Error(e.Cause))
			code = http.StatusUnauthorized
			msg = e.Error()
		case *customerrors.ForbiddenError:
			logger.Warn(e.DebugMessage, zap.Error(e.Cause))
			code = http.StatusForbidden
			msg = e.Error()
		case *customerrors.InternalServerError:
			logger.Error(e.DebugMessage, zap.Error(e.Cause))
			code = http.StatusInternalServerError
			msg = e.Error()
		}
	} else {
		switch err.(type) {
		case validator.ValidationErrors:
			logger.Warn("バリデーションエラー", zap.Error(err))
			code = http.StatusBadRequest
			msg = base.BadRequestError
		}
	}

	if code >= http.StatusInternalServerError {
		logger.Error(msg, zap.Int("status", code), zap.Error(err))
	} else {
		logger.Warn(msg, zap.Int("status", code), zap.Error(err))
	}

	// カスタムエラーレスポンスを送信
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			c.NoContent(code)
		} else {
			c.JSON(code, map[string]interface{}{
				"message": msg,
			})
		}
	}
}
