package usecase_study

import (
	"fmt"
	"github.com/k1e1n04/studios-api/base"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/auth"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	model_study "github.com/k1e1n04/studios-api/study/domain/model.study"
	repository_study "github.com/k1e1n04/studios-api/study/domain/repository.study"
)

// StudyUpdateService は 学習更新ユースケース
type StudyUpdateService struct {
	studyRepository repository_study.StudyRepository
	tagRepository   repository_study.TagRepository
}

// NewStudyUpdateService は 学習更新ユースケースを生成
func NewStudyUpdateService(
	studyRepository repository_study.StudyRepository,
	tagRepository repository_study.TagRepository,
) StudyUpdateService {
	return StudyUpdateService{
		studyRepository: studyRepository,
		tagRepository:   tagRepository,
	}
}

// Execute は学習更新を実行
func (sus *StudyUpdateService) Execute(param StudyUpdateParam) (*StudyDTO, error) {
	userID := *auth.RestoreUserID(param.UserID)
	targetStudy, err := sus.studyRepository.GetStudyByIDAndUserID(*model_study.RestoreStudyID(param.ID), userID)
	if err != nil {
		return nil, err
	}
	if targetStudy == nil {
		return nil, customerrors.NewBadRequestError(
			fmt.Sprintf("更新対象の学習が存在しません。ID: %s", param.ID),
			base.StudyNotFound,
			nil,
		)
	}
	var tags []*model_study.TagEntity
	if len(param.Tags) != 0 {
		tagEntities, err := sus.tagRepository.GetTagsByNamesAndUserID(param.Tags, userID)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tagEntities...)
		// tagが存在しない場合は作成する
		newTags := model_study.GenerateNotExistingTags(tagEntities, param.Tags, &userID)
		if len(newTags) != 0 {
			err = sus.tagRepository.CreateTags(newTags)
			if err != nil {
				return nil, err
			}
			tags = append(tags, newTags...)
		}
	}
	targetStudy.Update(param.Title, param.Content, tags)
	err = sus.studyRepository.UpdateStudy(targetStudy)
	if err != nil {
		return nil, err
	}
	return toStudyDTO(targetStudy), nil
}
