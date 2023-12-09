package study

import (
	"errors"
	"fmt"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	pagenation2 "github.com/k1e1n04/studios-api/base/sharedkarnel/model/pagenation"
	"github.com/k1e1n04/studios-api/base/usecase/pagenation"
	"github.com/k1e1n04/studios-api/src/adapter/infra/table"
	model_study "github.com/k1e1n04/studios-api/study/domain/model.study"
	repository_study "github.com/k1e1n04/studios-api/study/domain/repository.study"
	"gorm.io/gorm"
)

type StudyRepositoryImpl struct {
	db *gorm.DB
}

// NewStudyRepository は StudyRepository を生成
func NewStudyRepository(db *gorm.DB) repository_study.StudyRepository {
	return &StudyRepositoryImpl{
		db: db,
	}
}

// toStudyTableRecord は 学習テーブルレコードを生成
func toStudyTableRecord(study *model_study.StudyEntity) *table.Study {
	tags := make([]*table.Tag, len(study.Tags))
	for i, tag := range study.Tags {
		tags[i] = toTagTableRecord(tag)
	}
	return &table.Study{
		ID:          study.ID,
		Title:       study.Title,
		Content:     study.Content,
		Tags:        tags,
		CreatedDate: study.CreatedDate,
		UpdatedDate: study.UpdatedDate,
	}
}

// toStudyEntity は 学習テーブルレコードを学習エンティティに変換
func toStudyEntity(study *table.Study) *model_study.StudyEntity {
	tagEntities := make([]*model_study.TagEntity, len(study.Tags))
	for i, tag := range study.Tags {
		tagEntity := toTagEntity(tag)
		tagEntities[i] = tagEntity
	}
	return &model_study.StudyEntity{
		ID:          study.ID,
		Title:       study.Title,
		Content:     study.Content,
		Tags:        tagEntities,
		CreatedDate: study.CreatedDate,
		UpdatedDate: study.UpdatedDate,
	}
}

func toStudiesPage(studies []*table.Study, totalRecords int, pageable pagenation.Pageable) *model_study.StudiesPage {
	studyEntities := make([]*model_study.StudyEntity, len(studies))
	for i, study := range studies {
		studyEntity := toStudyEntity(study)
		studyEntities[i] = studyEntity
	}

	totalPages := totalRecords / pageable.Limit
	if totalRecords%pageable.Limit != 0 {
		totalPages++
	}

	return &model_study.StudiesPage{
		Page: pagenation2.Page{
			TotalElements: totalRecords,
			TotalPages:    totalPages,
			PageNumber:    pageable.Page,
			PageElements:  len(studies),
		},
		Studies: studyEntities,
	}
}

// CreateStudy はスタディを作成
func (r *StudyRepositoryImpl) CreateStudy(study *model_study.StudyEntity) error {
	studyTableRecord := toStudyTableRecord(study)
	err := r.db.Create(studyTableRecord).Error
	if err != nil {
		return customerrors.NewInternalServerError(
			fmt.Sprintf("学習の作成に失敗しました。 id: %s", study.ID),
			err,
		)
	}

	return nil
}

// UpdateStudy はスタディを更新
func (r *StudyRepositoryImpl) UpdateStudy(study *model_study.StudyEntity) error {
	studyTableRecord := toStudyTableRecord(study)
	err := r.db.Model(&studyTableRecord).Save(studyTableRecord).Error
	if err != nil {
		return customerrors.NewInternalServerError(
			fmt.Sprintf("学習の更新に失敗しました。 id: %s", study.ID),
			err,
		)
	}
	err = r.db.Debug().Model(&studyTableRecord).Association("Tags").Replace(&studyTableRecord.Tags)
	if err != nil {
		return customerrors.NewInternalServerError(
			fmt.Sprintf("学習のタグの更新に失敗しました。 id: %s", study.ID),
			err,
		)
	}
	return nil
}

// DeleteStudy はスタディを削除
func (r *StudyRepositoryImpl) DeleteStudy(study *model_study.StudyEntity) error {
	studyTableRecord := toStudyTableRecord(study)
	// アソシエーションを削除
	err := r.db.Model(&studyTableRecord).Association("Tags").Clear()
	if err != nil {
		return customerrors.NewInternalServerError(
			fmt.Sprintf("学習のタグの削除に失敗しました。 id: %s", study.ID),
			err,
		)
	}
	err = r.db.Delete(&studyTableRecord).Error
	if err != nil {
		return customerrors.NewInternalServerError(
			fmt.Sprintf("学習の削除に失敗しました。 id: %s", study.ID),
			err,
		)
	}
	return nil
}

// GetStudyByID はIDでスタディを取得
func (r *StudyRepositoryImpl) GetStudyByID(id string) (*model_study.StudyEntity, error) {
	var studyTableRecord table.Study
	err := r.db.Preload("Tags").First(&studyTableRecord, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, customerrors.NewInternalServerError(
				fmt.Sprintf("学習の取得に失敗しました。 id: %s", id),
				err,
			)
		}
	}
	study := toStudyEntity(&studyTableRecord)
	if err != nil {
		return nil, err
	}
	return study, nil
}

// GetStudiesByTitleOrTags はタイトルまたはタグからスタディを取得
func (r *StudyRepositoryImpl) GetStudiesByTitleOrTags(title string, tagName string, pageable pagenation.Pageable) (*model_study.StudiesPage, error) {
	var totalRecord int64
	var studies []*table.Study
	query := r.db.Preload("Tags").Table("studies").Model(&table.Study{}).Order("id DESC")

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	// タグ名で検索
	if tagName != "" {
		query = query.Joins("JOIN study_tags ON study_tags.study_id = studies.id")
		query = query.Joins("JOIN tags ON tags.id = study_tags.tag_id")
		query = query.Where("tags.name LIKE ?", "%"+tagName+"%")
	}

	if err := query.Count(&totalRecord).Error; err != nil {
		return nil, customerrors.NewInternalServerError(
			fmt.Sprintf("学習の総レコード数の取得に失敗しました。"),
			err,
		)
	}

	query = query.Offset(pageable.Offset()).Limit(pageable.Limit)
	if err := query.Find(&studies).Error; err != nil {
		return nil, customerrors.NewInternalServerError(
			fmt.Sprintf("学習の取得に失敗しました。"),
			err,
		)
	}

	return toStudiesPage(studies, int(totalRecord), pageable), nil
}
