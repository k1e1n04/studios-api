package usecase_study

import (
	"fmt"
	"github.com/k1e1n04/studios-api/base"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	repository_study "github.com/k1e1n04/studios-api/study/domain/repository.study"
)

// StudyDetailService は 学習詳細サービス
type StudyDetailService struct {
	studyRepository repository_study.StudyRepository
}

// NewStudyDetailService は 学習詳細サービスを生成
func NewStudyDetailService(studyRepository repository_study.StudyRepository) StudyDetailService {
	return StudyDetailService{
		studyRepository: studyRepository,
	}
}

// Get は学習の詳細を取得
func (sds *StudyDetailService) Get(id string) (*StudyDTO, error) {
	studyEntity, err := sds.studyRepository.GetStudyByID(id)
	if err != nil {
		return nil, err
	}
	if studyEntity == nil {
		return nil, customerrors.NewNotFoundError(
			fmt.Sprintf("学習が存在しませんでした。 %s", id),
			base.StudyNotFound,
			nil,
		)
	}
	return toStudyDTO(studyEntity), nil
}
