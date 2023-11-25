package usecase_study

import (
	model_study "github.com/k1e1n04/studios-api/study/domain/model.study"
	repository_study "github.com/k1e1n04/studios-api/study/domain/repository.study"
)

// StudiesPageService は 学習ページサービス
type StudiesPageService struct {
	studyRepository repository_study.StudyRepository
}

// NewStudiesPageService は StudiesPageService を生成
func NewStudiesPageService(studyRepository repository_study.StudyRepository) StudiesPageService {
	return StudiesPageService{
		studyRepository: studyRepository,
	}
}

// toStudiesPageDTO は 学習ページDTOに変換
func toStudiesPageDTO(studies []*model_study.StudyEntity, nextExclusiveStartKey string) *StudiesPageDTO {
	// 学習DTOのスライスを生成
	studyDTOs := make([]*StudyDTO, len(studies))
	for i, study := range studies {
		studyDTOs[i] = toStudyDTO(study)
	}

	return &StudiesPageDTO{
		Studies:          studyDTOs,
		LastEvaluatedKey: nextExclusiveStartKey,
		TotalCount:       len(studies),
	}
}

// Get は 学習ページを取得
func (sps *StudiesPageService) Get(title string, tags string, limit int, exclusiveStartKey string) (*StudiesPageDTO, error) {
	// 学習を取得
	studies, nextExclusiveStartKey, err := sps.studyRepository.GetStudiesByTitleOrTags(title, tags, limit, exclusiveStartKey)
	if err != nil {
		return nil, err
	}

	// 学習ページDTOを生成
	studiesPageDTO := toStudiesPageDTO(studies, nextExclusiveStartKey)

	return studiesPageDTO, nil
}
