package todos

import (
	"context"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTodo(ctx context.Context, req CreateTodoRequest) (*Todo, error) {
	todoModel := &Todo{
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UserId:      req.UserId,
	}
	result, err := s.repo.Create(ctx, todoModel)

	if err != nil {
		return nil, err
	}
	return result, nil
}
