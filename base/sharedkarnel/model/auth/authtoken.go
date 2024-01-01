package auth

// AuthToken は 認証トークン
type AuthToken struct {
	// AccessToken は アクセストークン
	AccessToken string
	// RefreshToken は リフレッシュトークン
	RefreshToken string
}
