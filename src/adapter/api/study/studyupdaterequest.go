package study

// StudyUpdateRequest は学習更新リクエスト
type StudyUpdateRequest struct {
	// Title はタイトル
	Title string `json:"title" form:"title" validate:"required"`
	// Content は 内容
	Content string `json:"content" form:"content" validate:"required"`
	// Tags は 複数のタグ名
	Tags []string `json:"tags" form:"tags" validate:"required"`
}
