package study

import (
	"github.com/k1e1n04/studios-api/base"
	pagenation2 "github.com/k1e1n04/studios-api/base/adapter/api/pagenation"
	"github.com/k1e1n04/studios-api/base/config"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	"github.com/k1e1n04/studios-api/base/usecase/pagenation"
	usecase_study "github.com/k1e1n04/studios-api/study/usecase.study"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// StudyController は Studyコントローラ
type StudyController struct {
	studyRegisterService       usecase_study.StudyRegisterService
	studiesPageService         usecase_study.StudiesPageService
	studyDetailService         usecase_study.StudyDetailService
	studyUpdateService         usecase_study.StudyUpdateService
	studyDeleteService         usecase_study.StudyDeleteService
	studyReviewCompleteService usecase_study.StudyReviewCompleteService
	studiesReviewPageService   usecase_study.StudiesReviewPageService
}

// NewStudyController は StudyController を生成
func NewStudyController(
	studyRegisterService usecase_study.StudyRegisterService,
	studiesPageService usecase_study.StudiesPageService,
	studyDetailService usecase_study.StudyDetailService,
	studyUpdateService usecase_study.StudyUpdateService,
	studyDeleteService usecase_study.StudyDeleteService,
	studyReviewCompleteService usecase_study.StudyReviewCompleteService,
	studiesReviewPageService usecase_study.StudiesReviewPageService,
) StudyController {
	return StudyController{
		studyRegisterService:       studyRegisterService,
		studiesPageService:         studiesPageService,
		studyDetailService:         studyDetailService,
		studyUpdateService:         studyUpdateService,
		studyDeleteService:         studyDeleteService,
		studyReviewCompleteService: studyReviewCompleteService,
		studiesReviewPageService:   studiesReviewPageService,
	}
}

// toTagResponse は TagResponse を生成
func toTagResponse(dto *usecase_study.TagDTO) *TagResponse {
	return &TagResponse{
		ID:   dto.ID,
		Name: dto.Name,
	}
}

// toStudyRegisterResponse は StudyRegisterResponse を生成
func toStudyRegisterResponse(dto *usecase_study.StudyDTO) *StudyRegisterResponse {
	// CreatedDate, UpdatedDate を time.Time から string に変換
	const customDateFormat = "2006-01-02"

	createdDateStr := dto.CreatedDate.Format(customDateFormat)
	updatedDateStr := dto.UpdatedDate.Format(customDateFormat)

	tags := make([]*TagResponse, len(dto.Tags))
	for i, tag := range dto.Tags {
		tags[i] = toTagResponse(tag)
	}

	return &StudyRegisterResponse{
		ID:             dto.ID,
		Title:          dto.Title,
		Tags:           tags,
		Content:        dto.Content,
		NumberOfReview: dto.NumberOfReview,
		CreatedDate:    createdDateStr,
		UpdatedDate:    updatedDateStr,
	}
}

// toStudyResponse は StudyResponse を生成
func toStudyResponse(dto *usecase_study.StudyDTO) *StudyResponse {
	// CreatedDate, UpdatedDate を time.Time から string に変換
	const customDateFormat = "2006-01-02"

	createdDateStr := dto.CreatedDate.Format(customDateFormat)
	updatedDateStr := dto.UpdatedDate.Format(customDateFormat)

	tags := make([]*TagResponse, len(dto.Tags))
	for i, tag := range dto.Tags {
		tags[i] = toTagResponse(tag)
	}
	return &StudyResponse{
		ID:             dto.ID,
		Title:          dto.Title,
		Tags:           tags,
		Content:        dto.Content,
		NumberOfReview: dto.NumberOfReview,
		CreatedDate:    createdDateStr,
		UpdatedDate:    updatedDateStr,
	}
}

// toStudiesPageResponse は StudiesPageResponse を生成
func toStudiesPageResponse(dto *usecase_study.StudiesPageDTO) *StudiesPageResponse {
	// 学習DTOのスライスを生成
	studyResponses := make([]*StudyResponse, len(dto.Studies))
	for i, study := range dto.Studies {
		studyResponses[i] = toStudyResponse(study)
	}

	return &StudiesPageResponse{
		Studies: studyResponses,
		Page: pagenation2.PageResponse{
			TotalElements: dto.Page.TotalElements,
			TotalPages:    dto.Page.TotalPages,
			PageNumber:    dto.Page.PageNumber,
			PageElements:  dto.Page.PageElements,
		},
	}
}

// Register は 学習を登録
func (sc *StudyController) Register(c echo.Context) error {
	// ロガーをコンテキストから取得
	logger := c.Get(config.LoggerKey).(*zap.Logger)
	// ユーザーIDをコンテキストから取得
	userID := c.Get(config.UserIDKey).(string)
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
		return customerrors.NewBadRequestError(
			"リクエストが不正です",
			base.BadRequestError,
			err,
		)
	}
	dto, err := sc.studyRegisterService.Execute(usecase_study.StudyRegisterParam{
		Title:   studyRegisterRequest.Title,
		Tags:    studyRegisterRequest.Tags,
		Content: studyRegisterRequest.Content,
		UserID:  userID,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, toStudyRegisterResponse(dto))
}

// GetStudies は 学習一覧ページを取得
func (sc *StudyController) GetStudies(c echo.Context) error {
	// ロガーをコンテキストから取得
	logger := c.Get(config.LoggerKey).(*zap.Logger)
	// ユーザーIDをコンテキストから取得
	userID := c.Get(config.UserIDKey).(string)
	// パラメータを取得
	title := c.QueryParam("title")
	tagName := c.QueryParam("tag")
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		logger.Warn("クエリパラメータが不正です。", zap.Error(err))
		return customerrors.NewBadRequestError(
			"クエリパラメータが不正です。",
			base.InvalidPageNumber,
			err,
		)
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		logger.Warn("クエリパラメータが不正です。", zap.Error(err))
		return customerrors.NewBadRequestError(
			"クエリパラメータが不正です。",
			base.InvalidLimit,
			err,
		)
	}
	if err != nil {
		logger.Warn("size の変換に失敗しました", zap.Error(err))
		return customerrors.NewBadRequestError(
			"size の変換に失敗しました",
			base.InvalidSize,
			err,
		)
	}
	dto, err := sc.studiesPageService.Get(
		usecase_study.StudiesPageParam{
			Title:   title,
			TagName: tagName,
			UserID:  userID,
		},
		*pagenation.NewPageable(pageInt, limitInt),
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, toStudiesPageResponse(dto))
}

// GetStudy は学習の詳細を取得
func (sc *StudyController) GetStudy(c echo.Context) error {
	// id をパラメータから取得
	id := c.Param("id")
	// ユーザーIDをコンテキストから取得
	userID := c.Get(config.UserIDKey).(string)
	studyDTO, err := sc.studyDetailService.Get(id, userID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, toStudyResponse(studyDTO))
}

// UpdateStudy は学習を更新
func (sc *StudyController) UpdateStudy(c echo.Context) error {
	// ロガーをコンテキストから取得
	logger := c.Get(config.LoggerKey).(*zap.Logger)
	// ユーザーIDをコンテキストから取得
	userID := c.Get(config.UserIDKey).(string)
	// id をパラメータから取得
	id := c.Param("id")
	var studyUpdateRequest StudyUpdateRequest
	if err := c.Bind(&studyUpdateRequest); err != nil {
		logger.Warn("リクエストが不正です。", zap.Error(err))
		return customerrors.NewBadRequestError(
			"リクエストのバインドに失敗しました",
			base.InvalidJSONError,
			err,
		)
	}
	// バリデーション
	if err := c.Validate(&studyUpdateRequest); err != nil {
		logger.Warn("リクエストが不正です", zap.Error(err))
		return customerrors.NewBadRequestError(
			"リクエストが不正です",
			base.BadRequestError,
			err,
		)
	}
	dto, err := sc.studyUpdateService.Execute(usecase_study.StudyUpdateParam{
		ID:      id,
		Title:   studyUpdateRequest.Title,
		Content: studyUpdateRequest.Content,
		UserID:  userID,
		Tags:    studyUpdateRequest.Tags,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, toStudyResponse(dto))
}

// DeleteStudy は学習を削除
func (sc *StudyController) DeleteStudy(c echo.Context) error {
	// id をパラメータから取得
	id := c.Param("id")
	// ユーザーIDをコンテキストから取得
	userID := c.Get(config.UserIDKey).(string)
	err := sc.studyDeleteService.Execute(id, userID)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

// CompleteReview は学習の復習を完了
func (sc *StudyController) CompleteReview(c echo.Context) error {
	// id をパラメータから取得
	id := c.Param("id")
	// ユーザーIDをコンテキストから取得
	userID := c.Get(config.UserIDKey).(string)
	err := sc.studyReviewCompleteService.Execute(id, userID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

// GetStudiesReview は学習の復習一覧を取得
func (sc *StudyController) GetStudiesReview(c echo.Context) error {
	// ロガーをコンテキストから取得
	logger := c.Get(config.LoggerKey).(*zap.Logger)
	// ユーザーIDをコンテキストから取得
	userID := c.Get(config.UserIDKey).(string)
	// パラメータを取得
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		logger.Warn("クエリパラメータが不正です。", zap.Error(err))
		return customerrors.NewBadRequestError(
			"クエリパラメータが不正です。",
			base.InvalidPageNumber,
			err,
		)
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		logger.Warn("クエリパラメータが不正です。", zap.Error(err))
		return customerrors.NewBadRequestError(
			"クエリパラメータが不正です。",
			base.InvalidLimit,
			err,
		)
	}
	dto, err := sc.studiesReviewPageService.Get(
		userID,
		*pagenation.NewPageable(pageInt, limitInt),
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, toStudiesPageResponse(dto))
}
