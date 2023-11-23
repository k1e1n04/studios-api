package model_study

import (
	"github.com/google/uuid"
	"time"
)

// StudyEntity は 学びエンティティ
type StudyEntity struct {
	ID          string
	Title       string
	Tags        string
	Content     string
	CreatedDate time.Time
	UpdatedDate time.Time
}

// NewStudyEntity は StudyEntity を生成
func NewStudyEntity(title, tags, content string) *StudyEntity {
	return &StudyEntity{
		ID:          uuid.New().String(),
		Title:       title,
		Tags:        tags,
		Content:     content,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
}

// Update は StudyEntity を更新
func (s *StudyEntity) Update(title, tags, content string) {
	s.Title = title
	s.Tags = tags
	s.Content = content
	s.UpdatedDate = time.Now()
}

// FromDynamoDB は DynamoDB のアイテムから StudyEntity を生成
func FromDynamoDB(dynamodbItem map[string]string) *StudyEntity {
	createdDate, _ := time.Parse(time.RFC3339, dynamodbItem["created_date"])
	updatedDate, _ := time.Parse(time.RFC3339, dynamodbItem["updated_date"])

	return &StudyEntity{
		ID:          dynamodbItem["id"],
		Title:       dynamodbItem["title"],
		Tags:        dynamodbItem["tags"],
		Content:     dynamodbItem["content"],
		CreatedDate: createdDate,
		UpdatedDate: updatedDate,
	}
}
