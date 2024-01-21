package study

import (
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	"github.com/k1e1n04/studios-api/base/usecase/pagenation"
	"github.com/k1e1n04/studios-api/src/adapter/infra/table"
	usecase_study "github.com/k1e1n04/studios-api/study/usecase.study"
	"gorm.io/gorm"
	"time"
)

// StudiesReviewQueryServiceImpl は 学習復習クエリサービスの実体
type StudiesReviewQueryServiceImpl struct {
	db *gorm.DB
}

// NewStudiesReviewQueryService は 学習復習クエリサービスを生成
func NewStudiesReviewQueryService(db *gorm.DB) usecase_study.StudiesReviewQueryService {
	return &StudiesReviewQueryServiceImpl{
		db: db,
	}
}

// Get は 学習復習ページを取得
func (srqs *StudiesReviewQueryServiceImpl) Get(pageable pagenation.Pageable) (*usecase_study.StudiesPageDTO, error) {
	var totalRecord int64
	var studiesReviewSetting table.StudiesReviewSetting
	if err := srqs.db.First(&studiesReviewSetting).Error; err != nil {
		return nil, customerrors.NewInternalServerError(
			"学習復習設定の取得に失敗しました。",
			err,
		)
	}
	currentDate := time.Now().Format("2006-01-02")
	var studies []*table.Study
	query := srqs.db.Preload("Tags").
		Model(&table.Study{}).
		Where("studies.user_id = ?", userID.Value).
		Order("number_of_review asc").
		Where("((DATE_ADD(updated_date, INTERVAL ? DAY) <= ?) and number_of_review = 0) or "+
			"((DATE_ADD(updated_date, INTERVAL ? DAY) <= ?) and number_of_review = 1) or "+
			"((DATE_ADD(updated_date, INTERVAL ? DAY) <= ?) and number_of_review = 2)",
			studiesReviewSetting.FirstReviewInterval, currentDate,
			studiesReviewSetting.SecondReviewInterval, currentDate,
			studiesReviewSetting.ThirdReviewInterval, currentDate)

	if err := query.Count(&totalRecord).Error; err != nil {
		return nil, customerrors.NewInternalServerError(
			"学習の総レコード数の取得に失敗しました。",
			err,
		)
	}

	query = query.Offset(pageable.Offset()).Limit(pageable.Limit)
	if err := query.Find(&studies).Error; err != nil {
		return nil, customerrors.NewInternalServerError(
			"学習の取得に失敗しました。",
			err,
		)
	}

	return toStudiesPageDTO(studies, int(totalRecord), pageable), nil
}

// toStudiesPageDTO は 学習ページDTOに変換
func toStudiesPageDTO(studies []*table.Study, totalRecord int, pageable pagenation.Pageable) *usecase_study.StudiesPageDTO {
	// 学習DTOのスライスを生成
	studyDTOs := make([]*usecase_study.StudyDTO, len(studies))
	for i, study := range studies {
		studyDTOs[i] = toStudyDTO(study)
	}

	totalPages := totalRecord / pageable.Limit
	if totalRecord%pageable.Limit != 0 {
		totalPages++
	}

	return &usecase_study.StudiesPageDTO{
		Studies: studyDTOs,
		Page: pagenation.PageDTO{
			TotalElements: totalRecord,
			TotalPages:    totalPages,
			PageElements:  len(studies),
			PageNumber:    pageable.Page,
		},
	}
}

// toStudyDTO は 学習テーブルレコードを学習DTOに変換
func toStudyDTO(study *table.Study) *usecase_study.StudyDTO {
	tags := make([]*usecase_study.TagDTO, len(study.Tags))
	for i, tag := range study.Tags {
		tags[i] = toTagDTO(tag)
	}
	return &usecase_study.StudyDTO{
		ID:             study.ID,
		Title:          study.Title,
		Tags:           tags,
		Content:        study.Content,
		NumberOfReview: study.NumberOfReview,
		CreatedDate:    study.CreatedDate,
		UpdatedDate:    study.UpdatedDate,
	}
}

// toTagDTO は TagEntity を TagDTO に変換
func toTagDTO(tag *table.Tag) *usecase_study.TagDTO {
	return &usecase_study.TagDTO{
		ID:   tag.ID,
		Name: tag.Name,
	}
}
