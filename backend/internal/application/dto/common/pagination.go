package common

// Pagination リストレスポンス用のページネーション情報
type Pagination struct {
	Total   int  `json:"total"`
	Limit   int  `json:"limit"`
	Offset  int  `json:"offset"`
	HasMore bool `json:"has_more"`
}

// NewPagination 新しいページネーションインスタンスを作成
func NewPagination(total, limit, offset int) *Pagination {
	return &Pagination{
		Total:   total,
		Limit:   limit,
		Offset:  offset,
		HasMore: total > offset+limit,
	}
}