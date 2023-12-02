package di

import (
	"fmt"
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
		logger.Error("DBのDI登録に失敗しました。", zap.Error(err))
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
