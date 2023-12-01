package repository_study

import model_study "github.com/k1e1n04/studios-api/study/domain/model.study"

// TagRepository は タグリポジトリ
type TagRepository interface {
	// GetTagsByIDs は タグIDからタグを取得
	GetTagsByIDs(tagIds []string) ([]*model_study.TagEntity, error)
	// GetTagByID は タグIDからタグを取得
	GetTagByID(tagId string) (*model_study.TagEntity, error)
	// GetTagsByNames は タグ名からタグを取得
	GetTagsByNames(names []string) ([]*model_study.TagEntity, error)
	// CreateTags は タグを作成
	CreateTags(tag []*model_study.TagEntity) error
	// DeleteTags は タグを削除
	DeleteTags(tag []*model_study.TagEntity) error
	// SearchTags は タグを検索
	SearchTags(name string) ([]*model_study.TagEntity, error)
}
