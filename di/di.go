package di

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/k1e1n04/studios-api/study"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"os"
)

// RegisterDependencies は DI登録を行う
func RegisterDependencies(c *dig.Container, logger *zap.Logger) error {
	err := registerDB(c, logger)
	if err != nil {
		return err
	}

	err = registerRepository(c, logger)
	if err != nil {
		return err
	}

	err = registerUseCase(c, logger)
	if err != nil {
		return err
	}

	err = registerController(c, logger)
	if err != nil {
		return err
	}

	return nil
}

// registerDB は DB をコンテナに登録
func registerDB(bc *dig.Container, logger *zap.Logger) error {
	err := bc.Provide(func() (*dynamodb.DynamoDB, error) {
		if os.Getenv("ENV") == "Local" {
			// DynamoDB Localに接続するためのセッションを設定
			sess, err := session.NewSession(&aws.Config{
				Endpoint: aws.String("http://localhost:8000"),
				Region:   aws.String("ap-northeast-1"),
			})
			if err != nil {
				return nil, err
			}
			return dynamodb.New(sess), nil
		} else {
			// 通常のAWSセッションを使用してDynamoDBに接続
			sess, err := session.NewSession()
			if err != nil {
				return nil, err
			}
			return dynamodb.New(sess), nil
		}
	})

	if err != nil {
		logger.Error("DynamoDBの登録に失敗しました", zap.Error(err))
		return err
	}
	return nil
}

// registerRepository は Repository をコンテナに登録
func registerRepository(c *dig.Container, logger *zap.Logger) error {
	err := study.RegisterRepositoryToContainer(c, logger)
	if err != nil {
		return err
	}

	return nil
}

// registerUseCase は UseCase をコンテナに登録
func registerUseCase(bc *dig.Container, logger *zap.Logger) error {
	err := study.RegisterUseCaseToContainer(bc, logger)
	if err != nil {
		return err
	}

	return nil
}

// registerController は Controller をコンテナに登録
func registerController(bc *dig.Container, logger *zap.Logger) error {
	err := study.RegisterControllerToContainer(bc, logger)
	if err != nil {
		return err
	}

	return nil
}
