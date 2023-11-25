package study

import "github.com/k1e1n04/studios-api/base/adapter/api/pagenation"

// StudiesPageResponse は 学習ページレスポンス
type StudiesPageResponse struct {
	// Page は ページ
	Page pagenation.PageResponse `json:"page"`
	// Studies は 学習
	Studies []*StudyResponse `json:"studies"`
}
