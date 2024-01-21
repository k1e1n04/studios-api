package usecase_study

import (
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/auth"
	"github.com/k1e1n04/studios-api/base/usecase/pagenation"
)

// StudiesReviewQueryService は 学習復習クエリサービスのインターフェース
type StudiesReviewQueryService interface {
	// Get は 学習復習を取得
	Get(userID auth.UserID, pageable pagenation.Pageable) (*StudiesPageDTO, error)
}
