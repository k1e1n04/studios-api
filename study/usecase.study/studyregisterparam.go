package usecase_study

// StudyRegisterParam は 学習登録パラメータ
type StudyRegisterParam struct {
	// Title は タイトル
	Title string
	// Content は 内容
	Content string
	// Tags は 複数のタグ名
	Tags []string
}
