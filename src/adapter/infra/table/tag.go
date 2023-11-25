package table

// Tag は タグのテーブルレコード
type Tag struct {
	// ID は ID
	ID string `gorm:"primaryKey"`
	// Name は 名前
	Name string
}
