package customerrors

import "net/http"

// BadRequestError 不正なリクエスト時に発生するエラー
type BadRequestError struct {
	DebugMessage string
	FrontMessage string
	Cause        error
}

// NewBadRequestError BadRequestErrorを生成
func NewBadRequestError(
	debugMessage string,
	frontMessage string,
	cause error,
) *BadRequestError {
	return &BadRequestError{
		DebugMessage: debugMessage,
		FrontMessage: frontMessage,
		Cause:        cause,
	}
}

// Error エラーメッセージを返す
func (e *BadRequestError) Error() string {
	return e.FrontMessage
}

// StatusCode ステータスコードを返す
func (e *BadRequestError) StatusCode() int {
	return http.StatusBadRequest
}
