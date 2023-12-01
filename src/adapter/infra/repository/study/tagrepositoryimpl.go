package study

import (
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
	return &table.Tag{
		ID:   tag.ID,
		Name: tag.Name,
	}
}

// toTagEntity は タグエンティティに変換
func toTagEntity(tag *table.Tag) *model_study.TagEntity {
	return &model_study.TagEntity{
		ID:   tag.ID,
		Name: tag.Name,
	}
}

// GetTagsByIDs は タグIDからタグを取得
func (tri *TagRepositoryImpl) GetTagsByIDs(tagIds []string) ([]*model_study.TagEntity, error) {
	var tags []*model_study.TagEntity
	var tagTableRecords []*table.Tag
	if err := tri.DB.Where("id IN ?", tagIds).Find(&tagTableRecords).Error; err != nil {
		return nil, err
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
		return nil, err
	}
	return toTagEntity(&tagTableRecord), nil
}

// GetTagsByNames は タグ名からタグを取得
func (tri *TagRepositoryImpl) GetTagsByNames(names []string) ([]*model_study.TagEntity, error) {
	var tags []*model_study.TagEntity
	var tagTableRecords []*table.Tag
	if err := tri.DB.Where("name IN ?", names).Find(&tagTableRecords).Error; err != nil {
		return nil, err
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
		return err
	}
	return nil
}
