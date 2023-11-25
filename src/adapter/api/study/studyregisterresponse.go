package study

import "time"

// StudyRegisterResponse は 学習登録レスポンス
type StudyRegisterResponse struct {
	// ID はID
	ID string `json:"id"`
	// Title は タイトル
	Title string `json:"title"`
	// Content は 内容
	Content string `json:"content"`
	// Tags は タグ
	Tags string `json:"tags"`
	// CreatedDate は 作成日
	CreatedDate time.Time `json:"created_date"`
	// UpdatedDate は 更新日
	UpdatedDate time.Time `json:"updated_date"`
}
