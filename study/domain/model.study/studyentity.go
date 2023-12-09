package model_study

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

// StudyEntity は 学びエンティティ
type StudyEntity struct {
	// ID は ID
	ID string
	// Title は タイトル
	Title string
	// Tags は タグ
	Tags []*TagEntity
	// Content は 内容
	Content string
	// NumberOfReview は 復習回数
	NumberOfReview int
	// CreatedDate は 作成日
	CreatedDate time.Time
	// UpdatedDate は 更新日
	UpdatedDate time.Time
}

// NewStudyEntity は StudyEntity を生成
func NewStudyEntity(title, content string, tags []*TagEntity) *StudyEntity {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return &StudyEntity{
		ID:             id.String(),
		Title:          title,
		Content:        content,
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
func (s *StudyEntity) GetTagIDs() []string {
	tagIDs := make([]string, len(s.Tags))
	for i, tag := range s.Tags {
		tagIDs[i] = tag.ID
	}
	return tagIDs
}

// IncrementNumberOfReview は 復習回数をインクリメント
func (s *StudyEntity) IncrementNumberOfReview() {
	s.NumberOfReview++
}
