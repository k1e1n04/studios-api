package study

import (
	"github.com/k1e1n04/studios-api/base/config"
	usecase_study "github.com/k1e1n04/studios-api/study/usecase.study"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// TagController は タグコントローラ
type TagController struct {
	tagsGetService usecase_study.TagsGetService
}

// toTagsResponse は タグをレスポンスに変換
func toTagsResponse(tags []*usecase_study.TagDTO) *TagsResponse {
	tagResponses := make([]*TagResponse, len(tags))
	for i, tag := range tags {
		tagResponses[i] = toTagResponse(tag)
	}
	return &TagsResponse{
		Tags: tagResponses,
	}
}

// NewTagController は タグコントローラを生成
func NewTagController(tagsGetService usecase_study.TagsGetService) TagController {
	return TagController{
		tagsGetService: tagsGetService,
	}
}

// GetTags は タグを取得
func (tc *TagController) GetTags(c echo.Context) error {
	// ロガーをコンテキストから取得
	logger := c.Get(config.LoggerKey).(*zap.Logger)
	// ユーザーIDをコンテキストから取得
	userID := c.Get(config.UserIDKey).(string)
	tagName := c.QueryParam("tag")
	tags, err := tc.tagsGetService.Execute(tagName, userID)
	if err != nil {
		logger.Info("タグの取得に失敗しました", zap.Error(err))
		return err
	}
	return c.JSON(200, toTagsResponse(tags))
}
