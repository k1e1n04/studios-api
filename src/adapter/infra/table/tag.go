package table

// Tag は タグのテーブルレコード
type Tag struct {
	// ID は ID
	ID string `gorm:"primaryKey"`
	// Name は 名前
	Name string `gorm:"not null"`
	// UserID は ユーザーID
	UserID string `gorm:"not null"`
	// Studies は 学習
	Studies []*Study `gorm:"many2many:study_tags;foreignKey:ID;references:ID"`
}
