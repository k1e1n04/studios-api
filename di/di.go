package di

import (
	repositoryExample "github.com/k1e1n04/studios-api/example/domain/repository.example"
	usecaseExample "github.com/k1e1n04/studios-api/example/usecase.example"
	controllerExample "github.com/k1e1n04/studios-api/src/adapter/api/example"
	repositoryExampleImpl "github.com/k1e1n04/studios-api/src/adapter/infra/repository/example"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

// RegisterDependencies は DI登録を行う
func RegisterDependencies(c *dig.Container, logger *zap.Logger) error {
	err := registerExample(c, logger)
	if err != nil {
		return err
	}

	return nil
}

func registerExample(bc *dig.Container, logger *zap.Logger) error {
	// Repository
	err := bc.Provide(func() repositoryExample.ExampleRepository {
		return repositoryExampleImpl.NewExampleRepository()
	})
	if err != nil {
		logger.Panic("ExampleRepositoryのDI登録に失敗しました。", zap.Error(err))
		return err
	}

	// UseCase
	err = bc.Provide(func(repositoryExample repositoryExample.ExampleRepository) usecaseExample.ExampleService {
		return usecaseExample.NewExampleService(repositoryExample)
	})
	if err != nil {
		logger.Panic("ExampleServiceのDI登録に失敗しました。", zap.Error(err))
		return err
	}

	// Controller
	err = bc.Provide(func(exampleService usecaseExample.ExampleService) controllerExample.ExampleController {
		return controllerExample.NewExampleController(exampleService)
	})
	if err != nil {
		logger.Panic("ExampleControllerのDI登録に失敗しました。", zap.Error(err))
		return err
	}

	return nil
}
