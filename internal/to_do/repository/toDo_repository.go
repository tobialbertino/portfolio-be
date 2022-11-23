package repository

import (
	"context"
	"tobialbertino/portfolio-be/internal/to_do/models/entity"

	"github.com/jackc/pgx/v5"
)

type ToDoRepository interface {
	Create(ctx context.Context, db *pgx.Conn, toDo *entity.ToDo) (int64, error)
	Update(ctx context.Context, db *pgx.Conn, toDo *entity.ToDo) (*entity.ToDo, error)
	Delete(ctx context.Context, db *pgx.Conn, toDo *entity.ToDo) error
	GetAll(ctx context.Context, db *pgx.Conn) (*[]entity.ToDo, error)
}