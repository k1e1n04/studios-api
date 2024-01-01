package cognito

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

// CognitoClientAPI は Cognito のクライアントのインターフェース
type CognitoClientAPI interface {
	SignUp(
		ctx context.Context,
		params *cognitoidentityprovider.SignUpInput,
		optFns ...func(*cognitoidentityprovider.Options),
	) (*cognitoidentityprovider.SignUpOutput, error)
	InitiateAuth(ctx context.Context,
		params *cognitoidentityprovider.InitiateAuthInput,
		optFns ...func(*cognitoidentityprovider.Options),
	) (*cognitoidentityprovider.InitiateAuthOutput, error)
	ListUsers(
		ctx context.Context,
		params *cognitoidentityprovider.ListUsersInput,
		optFns ...func(*cognitoidentityprovider.Options),
	) (*cognitoidentityprovider.ListUsersOutput, error)
}

// Cognito 型
type Cognito struct {
	Client     CognitoClientAPI
	UserPoolID string
	ClientID   string
}

// SignUpInput は Cognito へのサインアップ時の型
type SignUpInput struct {
	ClientId       string
	Email          string
	Username       string
	Password       string
	AgreeToTerms   bool
	UserAttributes []types.AttributeType
}

// NewCognito は Cognito を生成する関数
func NewCognito(client CognitoClientAPI) *Cognito {
	return &Cognito{
		Client:     client,
		UserPoolID: os.Getenv("COGNITO_USER_POOL_ID"),
		ClientID:   os.Getenv("COGNITO_CLIENT_ID"),
	}
}
