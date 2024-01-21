package auth

import (
	usecase_auth "github.com/k1e1n04/studios-api/auth/usecase.auth"
	"github.com/k1e1n04/studios-api/base"
	"github.com/k1e1n04/studios-api/base/config"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

// AuthController は 認証コントローラー
type AuthController struct {
	signUpService usecase_auth.SignUpService
	loginService  usecase_auth.LoginService
}

// NewAuthController は AuthController を生成
func NewAuthController(
	signUpService usecase_auth.SignUpService,
	loginService usecase_auth.LoginService,
) AuthController {
	return AuthController{
		signUpService: signUpService,
		loginService:  loginService,
	}
}

// SignUp は サインアップを実行
func (ac *AuthController) SignUp(c echo.Context) error {
	logger := c.Get(config.LoggerKey).(*zap.Logger)
	req := SignupRequest{}
	err := c.Bind(&req)
	if err != nil {
		logger.Warn("リクエストが不正です。", zap.Error(err))
		return customerrors.NewBadRequestError(
			"リクエストのバインドに失敗しました",
			base.InvalidJSONError,
			err,
		)
	}
	err = c.Validate(&req)
	if err != nil {
		logger.Warn("リクエストが不正です", zap.Error(err))
		return customerrors.NewBadRequestError(
			"リクエストが不正です",
			base.BadRequestError,
			err,
		)
	}
	if err := ac.signUpService.Execute(usecase_auth.SignUpParam{
		Email:           req.Email,
		Username:        req.Username,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
		AgreeToTerms:    req.AgreeToTerms,
	}); err != nil {
		logger.Info("サインアップに失敗しました", zap.Error(err))
		return err
	}
	return c.JSON(http.StatusOK, SignupResponse{
		Message: "サインアップに成功しました",
	})
}

// Login は ログインを実行
func (ac *AuthController) Login(c echo.Context) error {
	// ロガーをコンテキストから取得
	logger := c.Get(config.LoggerKey).(*zap.Logger)
	var req LoginRequest
	err := c.Bind(&req)
	if err != nil {
		logger.Warn("リクエストが不正です。", zap.Error(err))
		return customerrors.NewBadRequestError(
			"リクエストのバインドに失敗しました",
			base.InvalidJSONError,
			err,
		)
	}
	err = c.Validate(&req)
	if err != nil {
		logger.Warn("リクエストが不正です", zap.Error(err))
		return customerrors.NewBadRequestError(
			"リクエストが不正です",
			base.BadRequestError,
			err,
		)
	}
	dto, err := ac.loginService.Execute(usecase_auth.LoginParam{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		logger.Info("ログインに失敗しました", zap.Error(err))
		return err
	}
	return c.JSON(http.StatusOK, LoginResponse{
		AccessToken:  dto.AccessToken,
		RefreshToken: dto.RefreshToken,
	})
}
