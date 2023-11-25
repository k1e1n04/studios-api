package model_study

import "github.com/k1e1n04/studios-api/base/sharedkarnel/model/pagenation"

// StudiesPage は 学習のリストページ
type StudiesPage struct {
	Page    pagenation.Page
	Studies []*StudyEntity
}
