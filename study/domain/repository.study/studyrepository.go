package repository_study

import (
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/auth"
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
	// GetStudyByIDAndUserID は 学習IDとユーザーIDから学習を取得
	GetStudyByIDAndUserID(id model_study.StudyID, userID auth.UserID) (*model_study.StudyEntity, error)
	// GetStudiesByTitleOrTagsAndUserID は タイトルまたはタグとユーザーIDから学習を取得
	GetStudiesByTitleOrTagsAndUserID(
		title string, tagName string, userID auth.UserID, pageable pagenation.Pageable,
	) (*model_study.StudiesPage, error)
}
