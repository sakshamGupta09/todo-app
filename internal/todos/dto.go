package todos

import "time"

type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"required,min=3,max=100"`
	UserId      int    `json:"userId" validate:"required"`
}

type TodoResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserId      int       `json:"userId"`
}

type GetTodosRequest struct {
	PageNumber int `schema:"pageNumber"`
	PageSize   int `schema:"pageSize"`
}
