package usecase_auth

// LoginParam は ログインパラメータ
type LoginParam struct {
	// Email は メールアドレス
	Email string `json:"email"`
	// Password は パスワード
	Password string `json:"password"`
}
