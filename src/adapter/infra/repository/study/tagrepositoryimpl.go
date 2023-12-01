package study

import (
	"errors"
	"fmt"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	"github.com/k1e1n04/studios-api/src/adapter/infra/table"
	model_study "github.com/k1e1n04/studios-api/study/domain/model.study"
	repository_study "github.com/k1e1n04/studios-api/study/domain/repository.study"
	"gorm.io/gorm"
)

// TagRepositoryImpl は タグリポジトリ実体
type TagRepositoryImpl struct {
	// DB は DB
	DB *gorm.DB
}

// NewTagRepositoryImpl は タグリポジトリ実体を生成する
func NewTagRepositoryImpl(DB *gorm.DB) repository_study.TagRepository {
	return &TagRepositoryImpl{
		DB: DB,
	}
}

// toTagTableRecord は タグテーブルレコードに変換
func toTagTableRecord(tag *model_study.TagEntity) *table.Tag {
	studyTableRecords := make([]*table.Study, len(tag.Studies))
	for i, study := range tag.Studies {
		studyTableRecords[i] = toStudyTableRecord(study)
	}
	return &table.Tag{
		ID:      tag.ID,
		Name:    tag.Name,
		Studies: studyTableRecords,
	}
}

// toTagEntity は タグエンティティに変換
func toTagEntity(tag *table.Tag) *model_study.TagEntity {
	studyEntities := make([]*model_study.StudyEntity, len(tag.Studies))
	for i, study := range tag.Studies {
		studyEntities[i] = toStudyEntity(study)
	}
	return &model_study.TagEntity{
		ID:      tag.ID,
		Name:    tag.Name,
		Studies: studyEntities,
	}
}

// GetTagsByIDs は タグIDからタグを取得
func (tri *TagRepositoryImpl) GetTagsByIDs(tagIds []string) ([]*model_study.TagEntity, error) {
	var tags []*model_study.TagEntity
	var tagTableRecords []*table.Tag
	if err := tri.DB.Preload("Studies").Where("id IN ?", tagIds).Find(&tagTableRecords).Error; err != nil {
		return nil, customerrors.NewInternalServerError(
			fmt.Sprintf("タグの取得に失敗しました。 tagIds: %v", tagIds),
			err,
		)
	}
	for _, tagTableRecord := range tagTableRecords {
		tags = append(tags, toTagEntity(tagTableRecord))
	}
	return tags, nil
}

// GetTagByID は タグIDからタグを取得
func (tri *TagRepositoryImpl) GetTagByID(tagId string) (*model_study.TagEntity, error) {
	var tagTableRecord table.Tag
	if err := tri.DB.Where("id = ?", tagId).First(&tagTableRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, customerrors.NewInternalServerError(
			fmt.Sprintf("タグの取得に失敗しました。 tagId: %s", tagId),
			err,
		)
	}
	return toTagEntity(&tagTableRecord), nil
}

// GetTagsByNames は タグ名からタグを取得
func (tri *TagRepositoryImpl) GetTagsByNames(names []string) ([]*model_study.TagEntity, error) {
	var tags []*model_study.TagEntity
	var tagTableRecords []*table.Tag
	if err := tri.DB.Where("name IN ?", names).Find(&tagTableRecords).Error; err != nil {
		return nil, customerrors.NewInternalServerError(
			fmt.Sprintf("タグの取得に失敗しました。 names: %v", names),
			err,
		)
	}
	for _, tagTableRecord := range tagTableRecords {
		tags = append(tags, toTagEntity(tagTableRecord))
	}
	return tags, nil
}

// CreateTags は タグを作成
func (tri *TagRepositoryImpl) CreateTags(tag []*model_study.TagEntity) error {
	var tagTableRecords []*table.Tag
	for _, tagEntity := range tag {
		tagTableRecords = append(tagTableRecords, toTagTableRecord(tagEntity))
	}
	if err := tri.DB.Create(tagTableRecords).Error; err != nil {
		return customerrors.NewInternalServerError(
			fmt.Sprintf("タグの作成に失敗しました。 tag: %v", tag),
			err,
		)
	}
	return nil
}

// DeleteTags は タグを削除
func (tri *TagRepositoryImpl) DeleteTags(tag []*model_study.TagEntity) error {
	var tagTableRecords []*table.Tag
	for _, tagEntity := range tag {
		tagTableRecords = append(tagTableRecords, toTagTableRecord(tagEntity))
	}
	if err := tri.DB.Delete(tagTableRecords).Error; err != nil {
		return customerrors.NewInternalServerError(
			fmt.Sprintf("タグの削除に失敗しました。 tag: %v", tag),
			err,
		)
	}
	return nil
}
