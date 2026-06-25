package todos

import (
	"context"
	"errors"
	"todo-app/internal/errorCodes"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PaginatedResult struct {
	TotalRecords int
	Content      any
}

type PaginationParams struct {
	PageSize   int
	PageNumber int
}

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, todo *Todo) (*Todo, error) {
	query := `
		INSERT INTO todos (title, description, completed, createdAt, updatedAt, userId)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *
	`
	var insertedTodo Todo

	err := r.db.QueryRow(ctx, query, todo.Title, todo.Description, todo.Completed, todo.CreatedAt, todo.UpdatedAt, todo.UserId).Scan(&insertedTodo.Id, &insertedTodo.Title, &insertedTodo.Description, &insertedTodo.Completed, &insertedTodo.CreatedAt, &insertedTodo.UpdatedAt, &insertedTodo.UserId)

	if err != nil {
		return nil, errors.New(errorCodes.INTERNAL)
	}
	return &insertedTodo, nil
}

func (r *Repository) GetById(ctx context.Context, id int) (*Todo, error) {
	query := `
		SELECT
			id,
			title,
			description,
			completed,
			createdAt,
			updatedAt,
			userId
		FROM todos
		WHERE
			id = $1
	`

	var todo Todo

	err := r.db.QueryRow(ctx, query, id).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.UserId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New(errorCodes.NOT_FOUND)
		}
		return nil, errors.New(errorCodes.INTERNAL)
	}

	return &todo, nil
}

func (r *Repository) GetAll(ctx context.Context, userId int, params PaginationParams) (*PaginatedResult, error) {
	query := `
		SELECT
			COUNT(id)
		FROM todos
		WHERE
			user_id = $1
	`
	paginatedResult := &PaginatedResult{
		TotalRecords: 0,
		Content:      make([]Todo, 0),
	}
	err := r.db.QueryRow(ctx, query, userId).Scan(&paginatedResult.TotalRecords)
	if err != nil {
		return nil, errors.New(errorCodes.INTERNAL)
	}

	offset := (params.PageNumber - 1) * params.PageSize
	query = `
	SELECT
		id,
		title,
		description,
		completed,
		createdAt,
		updatedAt,
		userId
	FROM todos
	WHERE
		userId = $1
	OFFSET $2
	LIMIT $3
`
	rows, err := r.db.Query(ctx, query, userId, offset, params.PageSize)
	if err != nil {
		return nil, errors.New(errorCodes.INTERNAL)
	}
	defer rows.Close()

	todos, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Todo, error) {
		var todo Todo

		err = row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.UserId)
		return todo, err
	})

	if err != nil {
		return nil, errors.New(errorCodes.INTERNAL)
	}
	paginatedResult.Content = todos
	return paginatedResult, nil
}

func (r *Repository) Delete(ctx context.Context, todoId int) error {
	query := `
		DELETE FROM todos
		WHERE
			id = $1
	`
	result, err := r.db.Exec(ctx, query, todoId)

	if err != nil {
		return errors.New(errorCodes.INTERNAL)
	}

	if result.RowsAffected() == 0 {
		return errors.New(errorCodes.NOT_FOUND)
	}
	return nil
}
