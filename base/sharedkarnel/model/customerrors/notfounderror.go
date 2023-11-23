package customerrors

import "net/http"

// NotFoundError NotFound時に発生するエラー
type NotFoundError struct {
	DebugMessage string
	FrontMessage string
	Cause        error
}

// NewNotFoundError NotFoundErrorを生成
func NewNotFoundError(
	debugMessage string,
	frontMessage string,
	cause error,
) *NotFoundError {
	return &NotFoundError{
		DebugMessage: debugMessage,
		FrontMessage: frontMessage,
		Cause:        cause,
	}
}

// Error エラーメッセージを返す
func (e *NotFoundError) Error() string {
	return e.FrontMessage
}

// StatusCode ステータスコードを返す
func (e *NotFoundError) StatusCode() int {
	return http.StatusNotFound
}
