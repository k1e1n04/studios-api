package usecase_study

import "github.com/k1e1n04/studios-api/base/usecase/pagenation"

// StudiesPageDTO は 学習ページDTO
type StudiesPageDTO struct {
	Page pagenation.PageDTO
	// Studies は 学習
	Studies []*StudyDTO
}
