package table

// StudiesReviewSetting は 学習復習設定のテーブルレコード
type StudiesReviewSetting struct {
	// FirstReviewInterval は 初回復習間隔
	FirstReviewInterval int `gorm:"not null;"`
	// SecondReviewInterval は 2回目復習間隔
	SecondReviewInterval int `gorm:"not null;"`
	// ThirdReviewInterval は 3回目復習間隔
	ThirdReviewInterval int `gorm:"not null;"`
}

func (StudiesReviewSetting) TableName() string {
	return "studies_review_setting"
}
