package model_study

import (
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/auth"
	"time"
)

// StudyEntity は 学びエンティティ
type StudyEntity struct {
	// ID は ID
	ID *StudyID
	// Title は タイトル
	Title string
	// Tags は タグ
	Tags []*TagEntity
	// Content は 内容
	Content string
	// UserID は ユーザーID
	UserID *auth.UserID
	// NumberOfReview は 復習回数
	NumberOfReview int
	// CreatedDate は 作成日
	CreatedDate time.Time
	// UpdatedDate は 更新日
	UpdatedDate time.Time
}

// NewStudyEntity は StudyEntity を生成
func NewStudyEntity(title, content string, userID *auth.UserID, tags []*TagEntity) *StudyEntity {
	return &StudyEntity{
		ID:             NewStudyID(),
		Title:          title,
		Content:        content,
		UserID:         userID,
		Tags:           tags,
		NumberOfReview: 0,
		CreatedDate:    time.Now(),
		UpdatedDate:    time.Now(),
	}
}

// Update は StudyEntity を更新
func (s *StudyEntity) Update(title, content string, tags []*TagEntity) {
	s.Title = title
	s.Content = content
	s.Tags = tags
	s.UpdatedDate = time.Now()
}

// GetTagIDs は タグIDを取得
func (s *StudyEntity) GetTagIDs() []*TagID {
	tagIDs := make([]*TagID, len(s.Tags))
	for i, tag := range s.Tags {
		tagIDs[i] = tag.ID
	}
	return tagIDs
}

// IncrementNumberOfReview は 復習回数をインクリメント
func (s *StudyEntity) IncrementNumberOfReview() {
	s.NumberOfReview++
}
