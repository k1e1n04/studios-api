package model_study

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

// TagID は タグID
type TagID struct {
	// Value は 値
	Value string
}

// NewTagID は TagID を生成
func NewTagID() *TagID {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return &TagID{
		Value: ulid.MustNew(ulid.Timestamp(t), entropy).String(),
	}
}

// RestoreTagID は TagID を復元
func RestoreTagID(value string) *TagID {
	return &TagID{
		Value: value,
	}
}
