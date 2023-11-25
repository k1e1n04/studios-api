package usecase_study

import "time"

// StudyDTO は
type StudyDTO struct {
	// ID は ID
	ID string
	// Title は タイトル
	Title string
	// Tags は タグ
	Tags string
	// Content は 内容
	Content string
	// CreatedDate は 作成日時
	CreatedDate time.Time
	// UpdatedDate は 更新日時
	UpdatedDate time.Time
}
