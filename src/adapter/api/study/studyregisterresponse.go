package study

// StudyRegisterResponse は 学習登録レスポンス
type StudyRegisterResponse struct {
	// ID はID
	ID string `json:"id"`
	// Title は タイトル
	Title string `json:"title"`
	// Content は 内容
	Content string `json:"content"`
	// NumberOfReview は 復習回数
	NumberOfReview int `json:"number_of_review"`
	// Tags は タグ
	Tags []*TagResponse `json:"tags"`
	// CreatedDate は 作成日
	CreatedDate string `json:"created_date"`
	// UpdatedDate は 更新日
	UpdatedDate string `json:"updated_date"`
}
