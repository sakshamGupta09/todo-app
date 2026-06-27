package todos

import (
	"context"
	"errors"
	"net/http"
	"todo-app/internal/models"
	"todo-app/internal/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, todo *Todo) (*Todo, *models.AppError) {
	query := `
		INSERT INTO todos (title, description, completed, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *
	`
	var insertedTodo Todo

	err := r.db.QueryRow(ctx, query, todo.Title, todo.Description, todo.Completed, todo.CreatedAt, todo.UpdatedAt).Scan(&insertedTodo.Id, &insertedTodo.Title, &insertedTodo.Description, &insertedTodo.Completed, &insertedTodo.CreatedAt, &insertedTodo.UpdatedAt)

	if err != nil {
		return nil, utils.CreateError(http.StatusInternalServerError, INTERNAL_ERROR, err.Error())
	}
	return &insertedTodo, nil
}

func (r *Repository) GetById(ctx context.Context, id int) (*Todo, *models.AppError) {
	query := `
		SELECT
			id,
			title,
			description,
			completed,
			created_at,
			updated_at
		FROM todos
		WHERE
			id = $1
	`

	var todo Todo

	err := r.db.QueryRow(ctx, query, id).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.CreateError(http.StatusNotFound, NOT_FOUND, err.Error())
		}
		return nil, utils.CreateError(http.StatusInternalServerError, INTERNAL_ERROR, err.Error())
	}

	return &todo, nil
}

func (r *Repository) GetAll(ctx context.Context, params GetTodosRequest) (*models.PaginatedResponse[Todo], *models.AppError) {
	query := `
		SELECT
			COUNT(id)
		FROM todos
	`
	paginatedResult := &models.PaginatedResponse[Todo]{
		TotalRecords: 0,
		Content:      make([]Todo, 0),
		PageSize:     params.PageSize,
		PageNumber:   params.PageNumber,
	}
	err := r.db.QueryRow(ctx, query).Scan(&paginatedResult.TotalRecords)
	if err != nil {
		return nil, utils.CreateError(http.StatusInternalServerError, INTERNAL_ERROR, err.Error())
	}

	offset := (params.PageNumber - 1) * params.PageSize
	query = `
	SELECT
		id,
		title,
		description,
		completed,
		created_at,
		updated_at
	FROM todos
	OFFSET $1
	LIMIT $2
`
	rows, err := r.db.Query(ctx, query, offset, params.PageSize)
	if err != nil {
		return nil, utils.CreateError(http.StatusInternalServerError, INTERNAL_ERROR, err.Error())
	}
	defer rows.Close()

	todos, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Todo, error) {
		var todo Todo

		err = row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		return todo, err
	})

	if err != nil {
		return nil, utils.CreateError(http.StatusInternalServerError, INTERNAL_ERROR, err.Error())
	}
	paginatedResult.Content = todos
	return paginatedResult, nil
}

func (r *Repository) Delete(ctx context.Context, todoId int) *models.AppError {
	query := `
		DELETE FROM todos
		WHERE
			id = $1
	`
	result, err := r.db.Exec(ctx, query, todoId)

	if err != nil {
		return utils.CreateError(http.StatusInternalServerError, INTERNAL_ERROR, err.Error())
	}

	if result.RowsAffected() == 0 {
		return utils.CreateError(http.StatusNotFound, NOT_FOUND, NOT_FOUND)
	}
	return nil
}
