package repository_study

import (
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/auth"
	model_study "github.com/k1e1n04/studios-api/study/domain/model.study"
)

// TagRepository は タグリポジトリ
type TagRepository interface {
	// GetTagsByIDsAndUserID は タグIDとユーザーIDからタグを取得
	GetTagsByIDsAndUserID(tagIds []*model_study.TagID, userID auth.UserID) ([]*model_study.TagEntity, error)
	// GetTagByIDAndUserID は タグIDとユーザーIDからタグを取得
	GetTagByIDAndUserID(tagId *model_study.TagID, userID auth.UserID) (*model_study.TagEntity, error)
	// GetTagsByNamesAndUserID は タグ名とユーザーIDからタグを取得
	GetTagsByNamesAndUserID(names []string, userID auth.UserID) ([]*model_study.TagEntity, error)
	// CreateTags は タグを作成
	CreateTags(tag []*model_study.TagEntity) error
	// DeleteTags は タグを削除
	DeleteTags(tag []*model_study.TagEntity) error
	// SearchTags は タグを検索
	SearchTags(name string, userID auth.UserID) ([]*model_study.TagEntity, error)
}
