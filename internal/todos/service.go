package todos

import (
	"context"
	"time"
	"todo-app/internal/models"
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

	return result, err
}

func (s *Service) GetTodos(ctx context.Context, userId int, req GetTodosRequest) (*models.PaginatedResponse[Todo], error) {
	res, err := s.repo.GetAll(ctx, userId, req)
	return res, err
}

func (s *Service) GetTodoDetails(ctx context.Context, todoId int) (*Todo, error) {
	res, err := s.repo.GetById(ctx, todoId)
	return res, err
}

func (s *Service) DeleteTodo(ctx context.Context, todoId int) error {
	err := s.repo.Delete(ctx, todoId)
	return err
}
