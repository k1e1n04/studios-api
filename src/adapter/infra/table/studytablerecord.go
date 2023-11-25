package table

// StudyTableRecord は 学習のテーブルのレコード(DynamoDBのItem)の構造体
type StudyTableRecord struct {
	// ID は ID
	ID string `json:"id"`
	// Title は タイトル
	Title string `json:"title"`
	// Tags は タグ
	Tags string `json:"tags"`
	// Content は 内容
	Content string `json:"content"`
	// CreatedDate は 作成日時
	CreatedDate string `json:"createdDate"`
	// UpdatedDate は 更新日時
	UpdatedDate string `json:"updatedDate"`
}
