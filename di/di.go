package di

import (
	"github.com/k1e1n04/studios-api/study"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

// RegisterDependencies は DI登録を行う
func RegisterDependencies(c *dig.Container, logger *zap.Logger) error {
	err := registerRepository(c, logger)
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
