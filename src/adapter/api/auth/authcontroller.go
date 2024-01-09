package auth

import (
	usecase_auth "github.com/k1e1n04/studios-api/auth/usecase.auth"
	"github.com/labstack/echo/v4"
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
	req := SignUpRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}
	if err := ac.signUpService.Execute(usecase_auth.SignUpParam{
		Email:           req.Email,
		Username:        req.Username,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
		AgreeToTerms:    req.AgreeToTerms,
	}); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

// Login は ログインを実行
func (ac *AuthController) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}
	dto, err := ac.loginService.Execute(usecase_auth.LoginParam{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto)
}
