package usecase_auth

import (
	repository_auth "github.com/k1e1n04/studios-api/auth/domain/repository.auth"
	"github.com/k1e1n04/studios-api/base"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	"regexp"
	"strings"
)

// SignUpService は サインアップユースケース
type SignUpService struct {
	authRepository repository_auth.AuthRepository
}

// NewSignUpService は サインアップユースケースを生成
func NewSignUpService(authRepository repository_auth.AuthRepository) SignUpService {
	return SignUpService{
		authRepository: authRepository,
	}
}

// Execute は サインアップを実行
func (sus *SignUpService) Execute(param SignUpParam) error {
	if err := validateSignUpParam(param); err != nil {
		return err
	}
	return sus.authRepository.SignUp(param.Username, param.Email, param.Password)
}

// validateSignUpParam は サインアップパラメータのバリデーションを実行
func validateSignUpParam(param SignUpParam) error {
	if len(param.Email) == 0 {
		return customerrors.NewBadRequestError(
			"メールアドレスを入力してください",
			base.EmailRequired,
			nil,
		)
	}
	if !strings.Contains(param.Email, "@") {
		return customerrors.NewBadRequestError(
			"メールアドレスの形式が不正です",
			base.InvalidEmail,
			nil,
		)
	}
	if len(param.Username) < base.UsernameMinLength || len(param.Username) > base.UsernameMaxLength {
		return customerrors.NewBadRequestError(
			"ユーザー名は3文字以上20文字以下で入力してください",
			base.InvalidUsernameLength,
			nil,
		)
	}
	isAlphanumeric := regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
	if !isAlphanumeric(param.Username) {
		return customerrors.NewBadRequestError(
			"ユーザー名は半角英数字のみです",
			base.UsernameMustBeAlphanumeric,
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
	hasNumber := regexp.MustCompile(`\d`).MatchString
	if !hasNumber(param.Password) {
		return customerrors.NewBadRequestError(
			"パスワードには少なくとも1つの数字を含む必要があります",
			base.PasswordMustIncludeNumber,
			nil,
		)
	}
	hasSpecialCharacter := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString
	if !hasSpecialCharacter(param.Password) {
		return customerrors.NewBadRequestError(
			"パスワードには少なくとも1つの特殊文字を含む必要があります",
			base.PasswordMustIncludeSpecial,
			nil,
		)
	}
	hasUpperCase := regexp.MustCompile(`[A-Z]`).MatchString
	if !hasUpperCase(param.Password) {
		return customerrors.NewBadRequestError(
			"パスワードには少なくとも1つの大文字を含む必要があります",
			base.PasswordMustIncludeUpper,
			nil,
		)
	}
	hasLowerCase := regexp.MustCompile(`[a-z]`).MatchString
	if !hasLowerCase(param.Password) {
		return customerrors.NewBadRequestError(
			"パスワードには少なくとも1つの小文字を含む必要があります",
			base.PasswordMustIncludeLower,
			nil,
		)
	}
	if param.Password != param.PasswordConfirm {
		return customerrors.NewBadRequestError(
			"パスワードが一致しません",
			base.PasswordNotMatch,
			nil,
		)
	}
	if !param.AgreeToTerms {
		return customerrors.NewBadRequestError(
			"利用規約に同意していません",
			base.AgreeToTerms,
			nil,
		)
	}
	return nil
}
