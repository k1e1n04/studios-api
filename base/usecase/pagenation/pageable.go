package pagenation

// Pageable は ページネーションのパラメータ
type Pageable struct {
	Page  int
	Limit int
}

// NewPageable は Pageable を生成する
func NewPageable(page int, limit int) *Pageable {
	return &Pageable{
		Page:  page,
		Limit: limit,
	}
}

// Offset は ページネーションのオフセットを返す
func (p *Pageable) Offset() int {
	return (p.Page - 1) * p.Limit
}
