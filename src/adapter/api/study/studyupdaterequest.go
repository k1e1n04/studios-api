package study

// StudyUpdateRequest は学習更新リクエスト
type StudyUpdateRequest struct {
	// Title はタイトル
	Title string
	// Content は 内容
	Content string
	// Tags は 複数のタグ名
	Tags []string
}
