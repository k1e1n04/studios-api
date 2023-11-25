package study

// StudyResponse は 学習レスポンス
type StudyResponse struct {
	// ID は ID
	ID string `json:"id"`
	// Title は タイトル
	Title string `json:"title"`
	// Tags は タグ
	Tags []*TagResponse `json:"tags"`
	// Content は 内容
	Content string `json:"content"`
	// CreatedDate は 作成日
	CreatedDate string `json:"created_date"`
	// UpdatedDate は 更新日
	UpdatedDate string `json:"updated_date"`
}
