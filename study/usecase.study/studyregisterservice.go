package usecase_study

import (
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/auth"
	model_study "github.com/k1e1n04/studios-api/study/domain/model.study"
	repositoryStudy "github.com/k1e1n04/studios-api/study/domain/repository.study"
)

// StudyRegisterService は Studyサービス
type StudyRegisterService struct {
	studyRepository repositoryStudy.StudyRepository
	tagRepository   repositoryStudy.TagRepository
}

// NewStudyRegisterService は StudyRegisterService を生成
func NewStudyRegisterService(
	studyRepository repositoryStudy.StudyRepository,
	tagRepository repositoryStudy.TagRepository,
) StudyRegisterService {
	return StudyRegisterService{
		studyRepository: studyRepository,
		tagRepository:   tagRepository,
	}
}

// toStudyDTO は StudyEntity を StudyDTO に変換
func toStudyDTO(studyEntity *model_study.StudyEntity) *StudyDTO {
	tags := make([]*TagDTO, len(studyEntity.Tags))
	for i, tag := range studyEntity.Tags {
		tags[i] = toTagDTO(tag)
	}
	return &StudyDTO{
		ID:             studyEntity.ID.Value,
		Title:          studyEntity.Title,
		Tags:           tags,
		Content:        studyEntity.Content,
		NumberOfReview: studyEntity.NumberOfReview,
		CreatedDate:    studyEntity.CreatedDate,
		UpdatedDate:    studyEntity.UpdatedDate,
	}
}

// toTagDTO は TagEntity を TagDTO に変換
func toTagDTO(tagEntity *model_study.TagEntity) *TagDTO {
	return &TagDTO{
		ID:   tagEntity.ID.Value,
		Name: tagEntity.Name,
	}
}

// Execute は 学習を作成
func (srs *StudyRegisterService) Execute(param StudyRegisterParam) (*StudyDTO, error) {
	var tags []*model_study.TagEntity
	userID := *auth.RestoreUserID(param.UserID)
	var err error
	if len(param.Tags) != 0 {
		tags, err = srs.tagRepository.GetTagsByNamesAndUserID(param.Tags, userID)
		if err != nil {
			return nil, err
		}
		// 存在しないタグは新しく作成
		newTags := model_study.GenerateNotExistingTags(tags, param.Tags, &userID)
		if len(newTags) != 0 {
			err = srs.tagRepository.CreateTags(newTags)
			if err != nil {
				return nil, err
			}
			tags = append(tags, newTags...)
		}
	}
	studyEntity := model_study.NewStudyEntity(param.Title, param.Content, &userID, tags)
	err = srs.studyRepository.CreateStudy(studyEntity)
	if err != nil {
		return nil, err
	}
	return toStudyDTO(studyEntity), nil
}
