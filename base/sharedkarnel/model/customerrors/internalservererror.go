package customerrors

import "net/http"

// InternalServerError 内部エラー
type InternalServerError struct {
	DebugMessage string
	FrontMessage string
	Cause        error
}

// NewInternalServerError InternalServerErrorを生成
func NewInternalServerError(
	debugMessage string,
	cause error,
) *InternalServerError {
	return &InternalServerError{
		DebugMessage: debugMessage,
		FrontMessage: "サーバーエラーが発生しました。",
		Cause:        cause,
	}
}

// Error エラーメッセージを返す
func (e *InternalServerError) Error() string {
	return e.FrontMessage
}

// StatusCode ステータスコードを返す
func (e *InternalServerError) StatusCode() int {
	return http.StatusInternalServerError
}
