package pagenation

// PageResponse はページレスポンス
type PageResponse struct {
	// TotalElements は全要素数
	TotalElements int `json:"totalElements"`
	// TotalPages は全ページ数
	TotalPages int `json:"totalPages"`
	// PageNumber は現在のページ番号
	PageNumber int `json:"pageNumber"`
	// PageElements はページ要素数
	PageElements int `json:"pageElements"`
}
