package customerrors

import "net/http"

// ConflictError 競合時に発生するエラー
type ConflictError struct {
	DebugMessage string
	FrontMessage string
	Cause        error
}

// NewConflictError ConflictErrorを生成
func NewConflictError(
	debugMessage string,
	frontMessage string,
	cause error,
) *ConflictError {
	return &ConflictError{
		DebugMessage: debugMessage,
		FrontMessage: frontMessage,
		Cause:        cause,
	}
}

// Error エラーメッセージを返す
func (e *ConflictError) Error() string {
	return e.FrontMessage
}

// StatusCode ステータスコードを返す
func (e *ConflictError) StatusCode() int {
	return http.StatusConflict
}
