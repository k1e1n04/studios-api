package study

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	study2 "github.com/k1e1n04/studios-api/src/adapter/api/study"
	"github.com/k1e1n04/studios-api/src/adapter/infra/repository/study"
	repository_study "github.com/k1e1n04/studios-api/study/domain/repository.study"
	usecase_study "github.com/k1e1n04/studios-api/study/usecase.study"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

// RegisterRepositoryToContainer は Repository をコンテナに登録
func RegisterRepositoryToContainer(bc *dig.Container, logger *zap.Logger) error {
	err := bc.Provide(func(db *dynamodb.DynamoDB) repository_study.StudyRepository {
		return study.NewStudyRepository(db)
	})
	if err != nil {
		logger.Error("学習リポジトリの登録に失敗しました", zap.Error(err))
		return err
	}

	return nil
}

// RegisterUseCaseToContainer は UseCase をコンテナに登録
func RegisterUseCaseToContainer(bc *dig.Container, logger *zap.Logger) error {
	err := bc.Provide(func(studyRepository repository_study.StudyRepository) usecase_study.StudyRegisterService {
		return usecase_study.NewStudyRegisterService(studyRepository)
	})
	if err != nil {
		logger.Error("学習登録サービスの登録に失敗しました", zap.Error(err))
		return err
	}

	return nil
}

// RegisterControllerToContainer は Controller をコンテナに登録
func RegisterControllerToContainer(bc *dig.Container, logger *zap.Logger) error {
	err := bc.Provide(func(studyRegisterService usecase_study.StudyRegisterService) study2.StudyController {
		return study2.NewStudyController(studyRegisterService)
	})
	if err != nil {
		logger.Error("学習コントローラの登録に失敗しました", zap.Error(err))
		return err
	}

	return nil
}
