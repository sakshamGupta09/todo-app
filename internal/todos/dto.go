package todos

import "time"

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type GetTodosRequest struct {
	PageNumber int `schema:"pageNumber"`
	PageSize   int `schema:"pageSize"`
}
