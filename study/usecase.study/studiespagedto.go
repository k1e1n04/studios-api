package usecase_study

// StudiesPageDTO は 学習ページDTO
type StudiesPageDTO struct {
	// Studies は 学習
	Studies []*StudyDTO
	// LastEvaluatedKey は 最後に評価されたキー
	LastEvaluatedKey string
	// TotalCount は 総数
	TotalCount int
}
