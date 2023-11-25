package pagenation

// Page はページオブジェクト
type Page struct {
	// TotalElements は全要素数
	TotalElements int
	// TotalPages は全ページ数
	TotalPages int
	// PageElements はページ要素数
	PageElements int
	// PageNumber は現在のページ番号
	PageNumber int
}
