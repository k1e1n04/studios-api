package model_user

// UserEntity は ユーザーエンティティ
type UserEntity struct {
	// ID は ID
	ID string
	// Email は メールアドレス
	Email string
	// Username は ユーザー名
	Username string
	// agreeToTerms は 利用規約に同意したかどうか
	AgreeToTerms bool
}
