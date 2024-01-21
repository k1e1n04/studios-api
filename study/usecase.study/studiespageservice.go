package usecase_study

import (
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/auth"
	"github.com/k1e1n04/studios-api/base/usecase/pagenation"
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
func toStudiesPageDTO(studiesPage *model_study.StudiesPage) *StudiesPageDTO {
	// 学習DTOのスライスを生成
	studyDTOs := make([]*StudyDTO, len(studiesPage.Studies))
	for i, study := range studiesPage.Studies {
		studyDTOs[i] = toStudyDTO(study)
	}

	return &StudiesPageDTO{
		Studies: studyDTOs,
		Page: pagenation.PageDTO{
			TotalElements: studiesPage.Page.TotalElements,
			TotalPages:    studiesPage.Page.TotalPages,
			PageElements:  studiesPage.Page.PageElements,
			PageNumber:    studiesPage.Page.PageNumber,
		},
	}
}

// Get は 学習ページを取得
func (sps *StudiesPageService) Get(param StudiesPageParam, pageable pagenation.Pageable) (*StudiesPageDTO, error) {
	// 学習を取得
	entity, err := sps.studyRepository.GetStudiesByTitleOrTagsAndUserID(
		param.Title, param.TagName, *auth.RestoreUserID(param.UserID), pageable,
	)
	if err != nil {
		return nil, err
	}

	// 学習ページDTOを生成
	studiesPageDTO := toStudiesPageDTO(entity)

	return studiesPageDTO, nil
}
