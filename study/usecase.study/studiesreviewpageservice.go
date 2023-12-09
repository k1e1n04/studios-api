package usecase_study

import (
	"github.com/k1e1n04/studios-api/base/usecase/pagenation"
)

// StudiesReviewPageService は 学習復習ページサービス
type StudiesReviewPageService struct {
	studiesReviewQueryService StudiesReviewQueryService
}

// NewStudiesReviewPageService は StudiesReviewPageService を生成
func NewStudiesReviewPageService(studiesReviewQueryService StudiesReviewQueryService) StudiesReviewPageService {
	return StudiesReviewPageService{
		studiesReviewQueryService: studiesReviewQueryService,
	}
}

// Get は 学習復習ページを取得
func (srs *StudiesReviewPageService) Get(pageable pagenation.Pageable) (*StudiesPageDTO, error) {
	return srs.studiesReviewQueryService.Get(pageable)
}
