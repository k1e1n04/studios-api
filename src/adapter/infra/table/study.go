package table

import "time"

// Study は 学習のテーブルのテーブルレコード
type Study struct {
	// ID は ID
	ID string `gorm:"primaryKey;"`
	// Title は タイトル
	Title string `gorm:"not null;"`
	// Content は 内容
	Content string `gorm:"not null;"`
	// Tags は タグ
	Tags []*Tag `gorm:"many2many:study_tags;foreignKey:ID;references:ID"`
	// NumberOfReview は 復習回数
	NumberOfReview int `gorm:"not null;"`
	// CreatedDate は 作成日時
	CreatedDate time.Time `gorm:"not null;"`
	// UpdatedDate は 更新日時
	UpdatedDate time.Time `gorm:"not null;"`
}
