package usecase_study

import (
	model_study "github.com/k1e1n04/studios-api/study/domain/model.study"
	repositoryStudy "github.com/k1e1n04/studios-api/study/domain/repository.study"
)

// StudyRegisterService は Studyサービス
type StudyRegisterService struct {
	studyRepository repositoryStudy.StudyRepository
}

// NewStudyRegisterService は StudyRegisterService を生成
func NewStudyRegisterService(studyRepository repositoryStudy.StudyRepository) StudyRegisterService {
	return StudyRegisterService{
		studyRepository: studyRepository,
	}
}

// toStudyDTO は StudyEntity を StudyDTO に変換
func toStudyDTO(studyEntity *model_study.StudyEntity) *StudyDTO {
	return &StudyDTO{
		ID:          studyEntity.ID,
		Title:       studyEntity.Title,
		Tags:        studyEntity.Tags,
		Content:     studyEntity.Content,
		CreatedDate: studyEntity.CreatedDate,
		UpdatedDate: studyEntity.UpdatedDate,
	}
}

// Execute は 学習を作成
func (srs *StudyRegisterService) Execute(param StudyRegisterParam) (*StudyDTO, error) {
	studyEntity := model_study.NewStudyEntity(param.Title, param.Tags, param.Content)
	err := srs.studyRepository.CreateStudy(studyEntity)
	if err != nil {
		return nil, err
	}
	return toStudyDTO(studyEntity), nil
}
