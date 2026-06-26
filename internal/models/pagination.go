package models

type PaginatedResponse[T any] struct {
	TotalRecords int `json:"totalRecords"`
	Content      []T `json:"content"`
	PageSize     int `json:"pageSize"`
	PageNumber   int `json:"pageNumber"`
}

type PaginationParams struct {
	PageSize   int
	PageNumber int
}
