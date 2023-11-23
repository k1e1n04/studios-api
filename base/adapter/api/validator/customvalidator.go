package validator

import (
	"github.com/go-playground/validator/v10"
)

// CustomValidator はカスタムバリデーター
type CustomValidator struct {
	validator *validator.Validate
}

// NewCustomValidator はカスタムバリデーターを生成
func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

// Validate メソッドはバリデーションを実行します。
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
