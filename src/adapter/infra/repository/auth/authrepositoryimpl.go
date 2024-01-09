package auth

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	repository_auth "github.com/k1e1n04/studios-api/auth/domain/repository.auth"
	"github.com/k1e1n04/studios-api/base"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/auth"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	"github.com/k1e1n04/studios-api/src/adapter/infra/cognito"
	"strconv"
)

// AuthRepositoryImpl は 認証に関するリポジトリの実体
type AuthRepositoryImpl struct {
	cognito *cognito.Cognito
}

// NewAuthRepository は 認証に関するリポジトリを生成
func NewAuthRepository(cognito *cognito.Cognito) repository_auth.AuthRepository {
	return &AuthRepositoryImpl{
		cognito: cognito,
	}
}

// Login は ログイン処理を実行
func (ar *AuthRepositoryImpl) Login(email, password string) (*auth.AuthToken, error) {
	// Cognito を使用してユーザー認証を実行
	authResponse, err := ar.cognito.Client.InitiateAuth(context.TODO(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		AuthParameters: map[string]string{
			"USERNAME": email,
			"PASSWORD": password,
		},
		ClientId: aws.String(ar.cognito.ClientID),
	})
	if err != nil {
		var notAuthErr *types.NotAuthorizedException
		if errors.As(err, &notAuthErr) {
			return nil, customerrors.NewBadRequestError(
				"メールアドレスまたはパスワードが間違っています",
				base.InvalidEmailOrPassword,
				err,
			)
		}
		return nil, customerrors.NewInternalServerError(
			"ログインに失敗しました",
			err,
		)
	}
	return &auth.AuthToken{AccessToken: aws.ToString(authResponse.AuthenticationResult.AccessToken), RefreshToken: aws.ToString(authResponse.AuthenticationResult.RefreshToken)}, nil
}

// SignUp は サインアップ処理を実行
func (ar *AuthRepositoryImpl) SignUp(username, email, password string, agreeToTerms bool) error {
	// ユーザー登録処理
	var signUpInput *cognito.SignUpInput
	signUpInput = &cognito.SignUpInput{
		ClientId: ar.cognito.ClientID,
		Email:    email,
		Username: username,
		Password: password,
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("agreeToTerms"),
				Value: aws.String(strconv.FormatBool(agreeToTerms)),
			},
		},
	}
	_, err := ar.cognito.Client.SignUp(context.TODO(), &cognitoidentityprovider.SignUpInput{
		ClientId:       aws.String(signUpInput.ClientId),
		Username:       aws.String(signUpInput.Username),
		Password:       aws.String(signUpInput.Password),
		UserAttributes: signUpInput.UserAttributes,
	})
	if err != nil {
		var userExistsErr *types.UsernameExistsException
		var passErr *types.InvalidPasswordException
		var codeDeliveryErr *types.CodeDeliveryFailureException
		if errors.As(err, &userExistsErr) {
			return customerrors.NewBadRequestError(
				"ユーザーが既に存在しています",
				base.UserAlreadyExists,
				err,
			)
		} else if errors.As(err, &passErr) {
			return customerrors.NewBadRequestError(
				"パスワードが不正です",
				base.InvalidPassword,
				err,
			)
		} else if errors.As(err, &codeDeliveryErr) {
			return customerrors.NewInternalServerError(
				"メールの送信に失敗しました",
				err,
			)
		} else {
			return customerrors.NewInternalServerError(
				"ユーザー登録に失敗しました",
				err,
			)
		}
	}
	return err
}
