package table

// Tag は タグのテーブルレコード
type Tag struct {
	// ID は ID
	ID string `gorm:"primaryKey"`
	// Name は 名前
	Name string
	// Studies は 学習
	Studies []*Study `gorm:"many2many:study_tags;foreignKey:ID;references:ID"`
}
