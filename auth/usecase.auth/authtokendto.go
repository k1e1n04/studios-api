package usecase_auth

type AuthTokenDto struct {
	// AccessToken は アクセストークン
	AccessToken string
	// RefreshToken は リフレッシュトークン
	RefreshToken string
}
