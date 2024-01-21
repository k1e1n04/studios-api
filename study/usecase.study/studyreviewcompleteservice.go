package usecase_study

import (
	"fmt"
	"github.com/k1e1n04/studios-api/base"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/auth"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	model_study "github.com/k1e1n04/studios-api/study/domain/model.study"
	repository_study "github.com/k1e1n04/studios-api/study/domain/repository.study"
)

// StudyReviewCompleteService は 学習の復習完了サービス
type StudyReviewCompleteService struct {
	// studyRepository は 学習のリポジトリ
	studyRepository repository_study.StudyRepository
}

// NewStudyReviewCompleteService は StudyReviewCompleteService を生成
func NewStudyReviewCompleteService(studyRepository repository_study.StudyRepository) StudyReviewCompleteService {
	return StudyReviewCompleteService{
		studyRepository: studyRepository,
	}
}

// Execute は 学習の復習を完了
func (s *StudyReviewCompleteService) Execute(studyID string, userID string) error {
	study, err := s.studyRepository.GetStudyByIDAndUserID(*model_study.RestoreStudyID(studyID), *auth.RestoreUserID(userID))
	if err != nil {
		return err
	}
	if study == nil {
		return customerrors.NewNotFoundError(
			fmt.Sprintf("学習が存在しません。: %s", studyID),
			base.StudyNotFound,
			nil,
		)
	}
	study.IncrementNumberOfReview()
	if err := s.studyRepository.UpdateStudy(study); err != nil {
		return err
	}
	return nil
}
