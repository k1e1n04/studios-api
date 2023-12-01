package model_study

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

// StudyEntity は 学びエンティティ
type StudyEntity struct {
	ID          string
	Title       string
	Tags        []*TagEntity
	Content     string
	CreatedDate time.Time
	UpdatedDate time.Time
}

// NewStudyEntity は StudyEntity を生成
func NewStudyEntity(title, content string, tags []*TagEntity) *StudyEntity {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return &StudyEntity{
		ID:          id.String(),
		Title:       title,
		Content:     content,
		Tags:        tags,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
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
func (s *StudyEntity) GetTagIDs() []string {
	tagIDs := make([]string, len(s.Tags))
	for i, tag := range s.Tags {
		tagIDs[i] = tag.ID
	}
	return tagIDs
}
