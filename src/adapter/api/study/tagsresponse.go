package study

// TagsResponse は 複数のタグレスポンス
type TagsResponse struct {
	Tags []*TagResponse `json:"tags"`
}
