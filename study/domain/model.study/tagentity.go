package model_study

import (
	"github.com/oklog/ulid"
	"math/rand"
	"strings"
	"time"
)

// TagEntity は タグエンティティの構造体
type TagEntity struct {
	// ID は ID
	ID string
	// Name は 名前
	Name string
}

// NewTagEntity は タグエンティティを生成
func NewTagEntity(name string) *TagEntity {
	// TODO: IDの生成を共通化
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return &TagEntity{
		ID: id.String(),
		// 表記揺れを防ぐために小文字に変換
		Name: strings.ToLower(name),
	}
}

// GenerateNotExistingTags は 存在しないTagを生成
func GenerateNotExistingTags(existingTags []*TagEntity, tagNames []string) []*TagEntity {
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
			newTags = append(newTags, NewTagEntity(tagName))
		}
	}
	return newTags
}
