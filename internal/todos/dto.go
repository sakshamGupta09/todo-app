package todos

import "time"

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int    `json:"userId"`
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
