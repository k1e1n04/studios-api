package study

// StudyRegisterRequest は 学習登録リクエスト
type StudyRegisterRequest struct {
	// Title は タイトル
	Title string `json:"title" form:"title" validate:"required"`
	// Content は 内容
	Content string `json:"content" form:"content" validate:"required"`
	// Tags は タグ
	Tags []string `json:"tags" form:"tags" validate:"required"`
}
