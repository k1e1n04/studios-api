package customerrors

import "net/http"

// ForbiddenError アクセス権限がない時に発生するエラー
type ForbiddenError struct {
	DebugMessage string
	FrontMessage string
	Cause        error
}

// NewForbiddenError ForbiddenErrorを生成
func NewForbiddenError(
	debugMessage string,
	frontMessage string,
	cause error,
) *ForbiddenError {
	return &ForbiddenError{
		DebugMessage: debugMessage,
		FrontMessage: frontMessage,
		Cause:        cause,
	}
}

// Error エラーメッセージを返す
func (e *ForbiddenError) Error() string {
	return e.FrontMessage
}

// StatusCode ステータスコードを返す
func (e *ForbiddenError) StatusCode() int {
	return http.StatusForbidden
}
