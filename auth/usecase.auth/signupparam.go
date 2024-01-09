package usecase_auth

type SignUpParam struct {
	// Email は メールアドレス
	Email string `json:"email"`
	// Username は ユーザー名
	Username string `json:"username"`
	// Password は パスワード
	Password string `json:"password"`
	// PasswordConfirm は パスワード確認
	PasswordConfirm string `json:"password_confirm"`
	// AgreeToTerms は 利用規約に同意するかどうか
	AgreeToTerms bool `json:"agree_to_terms"`
}
