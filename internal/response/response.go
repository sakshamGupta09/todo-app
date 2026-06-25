package response

type PaginatedResponse[T any] struct {
	TotalRecords int   `json:"totalRecords"`
	Content      []T   `json:"content"`
	PageSize     int   `json:"pageSize"`
	PageNumber   int16 `json:"pageNumber"`
}

type PaginatedRequest struct {
	PageSize   int   `json:"pageSize"`
	PageNumber int16 `json:"pageNumber"`
}
