package user

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/k1e1n04/studios-api/base"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	"github.com/k1e1n04/studios-api/src/adapter/infra/cognito"
	model_user "github.com/k1e1n04/studios-api/user/domain/model.user"
	"os"
	"strconv"
)

// UserRepositoryImpl は ユーザーリポジトリの実体
type UserRepositoryImpl struct {
	cognito *cognito.Cognito
}

// NewUserRepository は ユーザーリポジトリを生成
func NewUserRepository(cognito *cognito.Cognito) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		cognito: cognito,
	}
}

// GetUserInfoByID は ユーザーIDからユーザー情報を取得する
func (ur *UserRepositoryImpl) GetUserInfoByID(id string) (*model_user.UserEntity, error) {
	resp, err := ur.cognito.Client.ListUsers(context.TODO(), &cognitoidentityprovider.ListUsersInput{
		UserPoolId: aws.String(os.Getenv("COGNITO_USER_POOL_ID")),
		Filter:     aws.String("sub=\"" + id + "\""),
	})
	if err != nil || len(resp.Users) == 0 {
		return nil, err
	}
	if len(resp.Users) == 0 {
		return nil, customerrors.NewNotFoundError(
			"ユーザーが見つかりませんでした",
			base.UserNotFound,
			err,
		)
	}
	user := resp.Users[0]
	// レスポンスからユーザー情報エンティティを作成する
	userEntity := &model_user.UserEntity{}
	for _, attr := range user.Attributes {
		switch aws.ToString(attr.Name) {
		case "sub":
			userEntity.ID = aws.ToString(attr.Value)
		case "username":
			userEntity.Username = aws.ToString(attr.Value)
		case "email":
			userEntity.Email = aws.ToString(attr.Value)
		case "agreeToTerms":
			userEntity.AgreeToTerms, err = strconv.ParseBool(aws.ToString(attr.Value))
			if err != nil {
				return nil, customerrors.NewInternalServerError(
					"ユーザー情報の取得に失敗しました",
					err,
				)
			}
		}
	}
	return userEntity, nil
}
