package auth

type LoginRequest struct {
	// Email は メールアドレス
	Email string `json:"email"  form:"email" validate:"required"`
	// Password は パスワード
	Password string `json:"password" form:"password" validate:"required"`
}
