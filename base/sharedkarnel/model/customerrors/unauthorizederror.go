package customerrors

import "net/http"

// UnauthorizedError 認証に失敗した時に発生するエラー
type UnauthorizedError struct {
	DebugMessage string
	FrontMessage string
	Cause        error
}

// NewUnauthorizedError UnauthorizedErrorを生成
func NewUnauthorizedError(
	debugMessage string,
	frontMessage string,
	cause error,
) *UnauthorizedError {
	return &UnauthorizedError{
		DebugMessage: debugMessage,
		FrontMessage: frontMessage,
		Cause:        cause,
	}
}

// Error エラーメッセージを返す
func (e *UnauthorizedError) Error() string {
	return e.FrontMessage
}

// StatusCode ステータスコードを返す
func (e *UnauthorizedError) StatusCode() int {
	return http.StatusUnauthorized
}
