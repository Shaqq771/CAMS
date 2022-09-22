package model

type Pagination struct {
	Limit     int `json:"limit_per_page"`
	Page      int `json:"current_page"`
	TotalPage int `json:"total_page"`
	TotalRows int `json:"total_rows"`
}

type QueryRequest struct {
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	SortBy string `query:"sort_by"`
	Search string `query:"search"`
}
