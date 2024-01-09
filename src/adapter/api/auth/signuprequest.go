package auth

type SignUpRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
	AgreeToTerms    bool   `json:"agree_to_terms" validate:"required"`
}
