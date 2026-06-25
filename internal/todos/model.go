package todos

import "time"

type Todo struct {
	Id          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserId      int
}
