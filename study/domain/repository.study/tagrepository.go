package repository_study

import model_study "github.com/k1e1n04/studios-api/study/domain/model.study"

// TagRepository は タグリポジトリ
type TagRepository interface {
	// GetTagsByIDs は タグIDからタグを取得
	GetTagsByIDs(tagIds []string) ([]*model_study.TagEntity, error)
	// GetTagByID は タグIDからタグを取得
	GetTagByID(tagId string) (*model_study.TagEntity, error)
	// CreateTag は タグを作成
	CreateTag(tag *model_study.TagEntity) error
}
