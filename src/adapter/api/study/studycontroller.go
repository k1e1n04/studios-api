package study

import (
	"github.com/k1e1n04/studios-api/base"
	"github.com/k1e1n04/studios-api/base/config"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	usecase_study "github.com/k1e1n04/studios-api/study/usecase.study"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

// StudyController は Studyコントローラ
type StudyController struct {
	studyRegisterService usecase_study.StudyRegisterService
}

// NewStudyController は StudyController を生成
func NewStudyController(studyRegisterService usecase_study.StudyRegisterService) StudyController {
	return StudyController{
		studyRegisterService: studyRegisterService,
	}
}

// toStudyRegisterResponse は StudyRegisterResponse を生成
func toStudyRegisterResponse(dto *usecase_study.StudyDTO) *StudyRegisterResponse {
	return &StudyRegisterResponse{
		ID:          dto.ID,
		Title:       dto.Title,
		Tags:        dto.Tags,
		Content:     dto.Content,
		CreatedDate: dto.CreatedDate,
		UpdatedDate: dto.UpdatedDate,
	}
}

// Register は 学習を登録
func (sc *StudyController) Register(c echo.Context) error {
	// ロガーをコンテキストから取得
	logger := c.Get(config.LoggerKey).(*zap.Logger)
	var studyRegisterRequest StudyRegisterRequest
	if err := c.Bind(&studyRegisterRequest); err != nil {
		logger.Warn("リクエストのバインドに失敗しました", zap.Error(err))
		return customerrors.NewBadRequestError(
			"リクエストのバインドに失敗しました",
			base.InvalidJSONError,
			err,
		)
	}
	// バリデーション
	if err := c.Validate(&studyRegisterRequest); err != nil {
		logger.Warn("リクエストが不正です", zap.Error(err))
		return err
	}
	dto, err := sc.studyRegisterService.Execute(usecase_study.StudyRegisterParam{
		Title:   studyRegisterRequest.Title,
		Tags:    studyRegisterRequest.Tags,
		Content: studyRegisterRequest.Content,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, toStudyRegisterResponse(dto))
}
