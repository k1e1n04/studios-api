package di

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/k1e1n04/studios-api/auth"
	"github.com/k1e1n04/studios-api/src/adapter/infra/cognito"
	"github.com/k1e1n04/studios-api/study"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// RegisterDependencies は DI登録を行う
func RegisterDependencies(c *dig.Container, logger *zap.Logger, dbUser string, dbPassword string) error {
	err := registerDB(c, logger, dbUser, dbPassword)
	if err != nil {
		return err
	}

	err = registerCognito(c, logger)
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
func registerDB(bc *dig.Container, logger *zap.Logger, dbUser string, dbPassword string) error {
	// DB
	err := bc.Provide(func() (*gorm.DB, error) {
		// MySQLに接続
		dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser,
			dbPassword,
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	})

	if err != nil {
		logger.Panic("DBのDI登録に失敗しました。", zap.Error(err))
		return err
	}
	return nil
}

// registerCognito は Cognito をコンテナに登録
func registerCognito(bc *dig.Container, logger *zap.Logger) error {
	// AWS SDKの設定
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		logger.Panic("unable to load SDK config, %v", zap.Error(err))
	}

	client := cognitoidentityprovider.NewFromConfig(cfg)

	co := cognito.NewCognito(client)

	err = bc.Provide(func() *cognito.Cognito {
		return co
	})
	if err != nil {
		logger.Panic("CognitoのDI登録に失敗しました。", zap.Error(err))
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

	err = auth.RegisterRepositoryToContainer(c, logger)
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

	err = auth.RegisterUseCaseToContainer(bc, logger)
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

	err = auth.RegisterControllerToContainer(bc, logger)
	if err != nil {
		return err
	}

	return nil
}
