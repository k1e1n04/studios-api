package auth

import (
	repository_auth "github.com/k1e1n04/studios-api/auth/domain/repository.auth"
	usecase_auth "github.com/k1e1n04/studios-api/auth/usecase.auth"
	auth2 "github.com/k1e1n04/studios-api/src/adapter/api/auth"
	"github.com/k1e1n04/studios-api/src/adapter/infra/cognito"
	"github.com/k1e1n04/studios-api/src/adapter/infra/repository/auth"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

// RegisterRepositoryToContainer は Repository をコンテナに登録
func RegisterRepositoryToContainer(bc *dig.Container, logger *zap.Logger) error {
	err := bc.Provide(func(cognito *cognito.Cognito) repository_auth.AuthRepository {
		return auth.NewAuthRepository(cognito)
	})
	if err != nil {
		logger.Error("認証リポジトリの登録に失敗しました", zap.Error(err))
		return err
	}

	return nil
}

// RegisterUseCaseToContainer は UseCase をコンテナに登録
func RegisterUseCaseToContainer(bc *dig.Container, logger *zap.Logger) error {
	err := bc.Provide(func(authRepository repository_auth.AuthRepository) usecase_auth.SignUpService {
		return usecase_auth.NewSignUpService(authRepository)
	})
	if err != nil {
		logger.Error("サインアップユースケースの登録に失敗しました", zap.Error(err))
		return err
	}

	err = bc.Provide(func(authRepository repository_auth.AuthRepository) usecase_auth.LoginService {
		return usecase_auth.NewLoginService(authRepository)
	})
	if err != nil {
		logger.Error("ログインユースケースの登録に失敗しました", zap.Error(err))
		return err
	}

	return nil
}

// RegisterControllerToContainer は Controller をコンテナに登録
func RegisterControllerToContainer(bc *dig.Container, logger *zap.Logger) error {
	err := bc.Provide(func(
		signUpService usecase_auth.SignUpService,
		loginService usecase_auth.LoginService,
	) auth2.AuthController {
		return auth2.NewAuthController(signUpService, loginService)
	})
	if err != nil {
		logger.Error("サインアップコントローラの登録に失敗しました", zap.Error(err))
		return err
	}

	return nil
}
