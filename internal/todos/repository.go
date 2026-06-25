package todos

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, todo *Todo) (int, error) {
	query := `
		INSERT INTO todos (title, description, completed, createdAt, updatedAt, userId)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	var insertedTodoId int

	err := r.db.QueryRow(ctx, query, todo.Title, todo.Description, todo.Completed, todo.CreatedAt, todo.UpdatedAt, todo.UserId).Scan(&insertedTodoId)

	if err != nil {
		return 0, err
	}
	return insertedTodoId, nil
}

func (r *Repository) GetById(ctx context.Context, id int) (Todo, error) {
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
			return Todo{}, errors.New("Todo not found")
		}
		return Todo{}, fmt.Errorf("%w", err)
	}

	return todo, nil
}

func (r *Repository) GetAll(ctx context.Context, userId int) ([]Todo, error) {
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
		userId = $1
`

	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer rows.Close()

	todos, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Todo, error) {
		var todo Todo

		err = row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.UserId)
		return todo, err
	})

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return todos, nil
}

func (r *Repository) Delete(ctx context.Context, todoId int) error {
	query := `
		DELETE FROM todos
		WHERE
			id = $1
	`
	result, err := r.db.Exec(ctx, query, todoId)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("Todo not found")
	}
	return nil
}
