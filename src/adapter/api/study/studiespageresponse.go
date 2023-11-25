package study

// StudiesPageResponse は 学習ページレスポンス
type StudiesPageResponse struct {
	// Studies は 学習
	Studies []*StudyResponse `json:"studies"`
	// LastEvaluatedKey は 最後に評価されたキー
	LastEvaluatedKey string `json:"lastEvaluatedKey"`
	// TotalCount は 総数
	TotalCount int `json:"totalCount"`
}
