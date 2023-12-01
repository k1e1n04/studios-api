package usecase_study

import (
	"fmt"
	"github.com/k1e1n04/studios-api/base"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	model_study "github.com/k1e1n04/studios-api/study/domain/model.study"
	repository_study "github.com/k1e1n04/studios-api/study/domain/repository.study"
)

// StudyDeleteService は 学習削除サービス
type StudyDeleteService struct {
	studyRepository repository_study.StudyRepository
	tagRepository   repository_study.TagRepository
}

// NewStudyDeleteService は 学習削除サービスを生成
func NewStudyDeleteService(
	studyRepository repository_study.StudyRepository,
	tagRepository repository_study.TagRepository,
) StudyDeleteService {
	return StudyDeleteService{
		studyRepository: studyRepository,
		tagRepository:   tagRepository,
	}
}

// Execute は 学習を削除
func (sds *StudyDeleteService) Execute(id string) error {
	targetStudy, err := sds.studyRepository.GetStudyByID(id)
	if err != nil {
		return err
	}
	if targetStudy == nil {
		return customerrors.NewBadRequestError(
			fmt.Sprintf("削除対象の学習が存在しません。ID: %s", id),
			base.StudyNotFound,
			nil,
		)
	}
	// 削除対象の学習以外に紐づいていないタグを削除
	relatedTags, err := sds.tagRepository.GetTagsByIDs(targetStudy.GetTagIDs())
	if err != nil {
		return err
	}
	err = sds.tagRepository.DeleteTags(getNotRelatedTags(relatedTags, targetStudy))
	if err != nil {
		return err
	}
	// 学習を削除
	err = sds.studyRepository.DeleteStudy(targetStudy)
	if err != nil {
		return err
	}
	return nil
}

// getNotRelatedTags は 削除対象の学習以外に紐づいていないタグを取得
func getNotRelatedTags(tags []*model_study.TagEntity, targetStudy *model_study.StudyEntity) []*model_study.TagEntity {
	var notRelatedTags []*model_study.TagEntity
	for _, tag := range tags {
		if len(tag.Studies) == 1 && tag.Studies[0].ID == targetStudy.ID {
			notRelatedTags = append(notRelatedTags, tag)
		}
	}
	return notRelatedTags
}
