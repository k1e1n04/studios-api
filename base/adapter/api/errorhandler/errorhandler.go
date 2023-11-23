package errorhandler

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/togisuma/standard-echo-serverless/base"
	"github.com/togisuma/standard-echo-serverless/base/config"
	"github.com/togisuma/standard-echo-serverless/base/sharedkarnel/model/customerrors"
	"go.uber.org/zap"
	"net/http"
)

// HTTPErrorHandler は HTTPエラーハンドラー
func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := base.InternalServerError

	var e interface{}

	// HTTPErrorのチェック
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
	} else if errors.As(err, &e) {
		switch e := e.(type) {
		case *customerrors.BadRequestError:
			code = http.StatusBadRequest
			msg = e.Error()
		case *customerrors.NotFoundError:
			code = http.StatusNotFound
			msg = e.Error()
		case *customerrors.ConflictError:
			code = http.StatusConflict
			msg = e.Error()
		case *customerrors.UnauthorizedError:
			code = http.StatusUnauthorized
			msg = e.Error()
		case *customerrors.ForbiddenError:
			code = http.StatusForbidden
			msg = e.Error()
		case *customerrors.InternalServerError:
			code = http.StatusInternalServerError
			msg = e.Error()
		}
	} else {
		switch err.(type) {
		case validator.ValidationErrors:
			code = http.StatusBadRequest
			msg = base.BadRequestError
		}
	}

	// エラーログを出力
	logger := c.Get(config.LoggerKey).(*zap.Logger)
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
