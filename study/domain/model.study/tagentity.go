package model_study

import (
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/auth"
	"strings"
)

// TagEntity は タグエンティティの構造体
type TagEntity struct {
	// ID は ID
	ID *TagID
	// Name は 名前
	Name string
	// UserID は ユーザーID
	UserID *auth.UserID
	// Studies は 学習
	Studies []*StudyEntity
}

// NewTagEntity は タグエンティティを生成
func NewTagEntity(name string, userID *auth.UserID) *TagEntity {
	// TODO: IDの生成を共通化
	return &TagEntity{
		ID: NewTagID(),
		// 表記揺れを防ぐために小文字に変換
		Name:   strings.ToLower(name),
		UserID: userID,
	}
}

// GenerateNotExistingTags は 存在しないTagを生成
func GenerateNotExistingTags(existingTags []*TagEntity, tagNames []string, userID *auth.UserID) []*TagEntity {
	newTags := make([]*TagEntity, 0)
	for _, tagName := range tagNames {
		isExist := false
		for _, tagEntity := range existingTags {
			if tagName == tagEntity.Name {
				isExist = true
				break
			}
		}
		if !isExist {
			newTags = append(newTags, NewTagEntity(tagName, userID))
		}
	}
	return newTags
}
