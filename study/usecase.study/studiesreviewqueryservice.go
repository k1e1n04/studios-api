package usecase_study

import "github.com/k1e1n04/studios-api/base/usecase/pagenation"

// StudiesReviewQueryService は 学習復習クエリサービスのインターフェース
type StudiesReviewQueryService interface {
	// Get は 学習復習を取得
	Get(pageable pagenation.Pageable) (*StudiesPageDTO, error)
}
