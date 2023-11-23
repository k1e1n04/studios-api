package model

// ErrorResponse エラー発生時のレスポンス
type ErrorResponse struct {
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}
