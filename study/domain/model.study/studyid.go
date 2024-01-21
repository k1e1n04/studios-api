package model_study

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

// StudyID は 学習ID
type StudyID struct {
	// Value は 値
	Value string
}

// NewStudyID は StudyID を生成
func NewStudyID() *StudyID {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return &StudyID{
		Value: ulid.MustNew(ulid.Timestamp(t), entropy).String(),
	}
}

// RestoreStudyID は StudyID を復元
func RestoreStudyID(value string) *StudyID {
	return &StudyID{
		Value: value,
	}
}
