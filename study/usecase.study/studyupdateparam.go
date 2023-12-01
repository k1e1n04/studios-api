package usecase_study

// StudyUpdateParam は学習更新パラメーター
type StudyUpdateParam struct {
	// ID は ID
	ID string
	// Title は タイトル
	Title string
	// Content は 内容
	Content string
	// Tags は 複数のタグ名
	Tags []string
}
