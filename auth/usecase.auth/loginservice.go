package usecase_auth

import (
	repository_auth "github.com/k1e1n04/studios-api/auth/domain/repository.auth"
	"github.com/k1e1n04/studios-api/base"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
)

// LoginService は ログインサービス
type LoginService struct {
	authRepository repository_auth.AuthRepository
}

// NewLoginService は ログインサービスを生成
func NewLoginService(authRepository repository_auth.AuthRepository) LoginService {
	return LoginService{
		authRepository: authRepository,
	}
}

// Execute は ログインを実行
func (ls *LoginService) Execute(param LoginParam) (*AuthTokenDto, error) {
	if err := validateLoginParam(param); err != nil {
		return nil, err
	}
	tokens, err := ls.authRepository.Login(param.Email, param.Password)
	if err != nil {
		return nil, err
	}
	return &AuthTokenDto{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

// validateLoginParam は ログインパラメータのバリデーションを実行
func validateLoginParam(param LoginParam) error {
	if len(param.Email) == 0 {
		return customerrors.NewBadRequestError(
			"メールアドレスが入力されていません",
			base.InvalidEmailOrPassword,
			nil,
		)
	}
	if len(param.Password) == 0 {
		return customerrors.NewBadRequestError(
			"パスワードが入力されていません",
			base.InvalidEmailOrPassword,
			nil,
		)
	}
	if len(param.Password) < base.PasswordMinLength {
		return customerrors.NewBadRequestError(
			"パスワードは8文字以上で入力してください",
			base.TooShortPassword,
			nil,
		)
	}
	return nil
}
