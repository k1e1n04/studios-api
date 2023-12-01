package repository_study

import (
	"github.com/k1e1n04/studios-api/base/usecase/pagenation"
	model_study "github.com/k1e1n04/studios-api/study/domain/model.study"
)

type StudyRepository interface {
	// CreateStudy は 学習を作成
	CreateStudy(study *model_study.StudyEntity) error
	// UpdateStudy は 学習を更新
	UpdateStudy(study *model_study.StudyEntity) error
	// DeleteStudy は 学習を削除
	DeleteStudy(study *model_study.StudyEntity) error
	// GetStudyByID は 学習を取得
	GetStudyByID(id string) (*model_study.StudyEntity, error)
	// GetStudiesByTitleOrTags は タイトルまたはタグから学習を取得
	GetStudiesByTitleOrTags(title string, tagName string, pageable pagenation.Pageable) (*model_study.StudiesPage, error)
}
