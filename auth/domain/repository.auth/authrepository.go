package repository_auth

import "github.com/k1e1n04/studios-api/base/sharedkarnel/model/auth"

// AuthRepository は 認証に関するリポジトリのインターフェース
type AuthRepository interface {
	// Login はログインを実行
	Login(email, password string) (*auth.AuthToken, error)
	// SignUp はサインアップを実行
	SignUp(username, email, password string, agreeToTerms bool) error
}
