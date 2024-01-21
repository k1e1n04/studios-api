package study

import (
	"errors"
	"fmt"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/auth"
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
		ID:      tag.ID.Value,
		Name:    tag.Name,
		UserID:  tag.UserID.Value,
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
		ID:      model_study.RestoreTagID(tag.ID),
		UserID:  auth.RestoreUserID(tag.UserID),
		Name:    tag.Name,
		Studies: studyEntities,
	}
}

// GetTagsByIDsAndUserID は タグIDとユーザーIDからタグを取得
func (tri *TagRepositoryImpl) GetTagsByIDsAndUserID(tagIds []*model_study.TagID, userID auth.UserID) ([]*model_study.TagEntity, error) {
	var tags []*model_study.TagEntity
	var tagTableRecords []*table.Tag
	tagIdStrings := make([]string, len(tagIds))
	for i, tagId := range tagIds {
		tagIdStrings[i] = tagId.Value
	}
	if err := tri.DB.Preload("Studies").Where("id IN ?", tagIdStrings, "user_id = ?", userID.Value).Find(&tagTableRecords).Error; err != nil {
		return nil, customerrors.NewInternalServerError(
			fmt.Sprintf("タグの取得に失敗しました。 tagIds: %v", tagIdStrings),
			err,
		)
	}
	for _, tagTableRecord := range tagTableRecords {
		tags = append(tags, toTagEntity(tagTableRecord))
	}
	return tags, nil
}

// GetTagByIDAndUserID は タグIDとユーザーIDからタグを取得
func (tri *TagRepositoryImpl) GetTagByIDAndUserID(tagId *model_study.TagID, userID auth.UserID) (*model_study.TagEntity, error) {
	var tagTableRecord table.Tag
	if err := tri.DB.Where("id = ?", tagId.Value, "user_id = ?", userID.Value).First(&tagTableRecord).Error; err != nil {
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

// GetTagsByNamesAndUserID は タグ名とユーザーIDからタグを取得
func (tri *TagRepositoryImpl) GetTagsByNamesAndUserID(names []string, userID auth.UserID) ([]*model_study.TagEntity, error) {
	var tags []*model_study.TagEntity
	var tagTableRecords []*table.Tag
	if err := tri.DB.Where("name IN ?", names, "user_id = ?", userID.Value).Find(&tagTableRecords).Error; err != nil {
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

// SearchTags は タグを検索
func (tri *TagRepositoryImpl) SearchTags(name string, userID auth.UserID) ([]*model_study.TagEntity, error) {
	var tags []*model_study.TagEntity
	var tagTableRecords []*table.Tag
	if err := tri.DB.Where("name LIKE ? AND user_id = ?", "%"+name+"%", userID.Value).Find(&tagTableRecords).Error; err != nil {
		return nil, customerrors.NewInternalServerError(
			fmt.Sprintf("タグの検索に失敗しました。 name: %s", name),
			err,
		)
	}
	for _, tagTableRecord := range tagTableRecords {
		tags = append(tags, toTagEntity(tagTableRecord))
	}
	return tags, nil
}
