package repository

import (
	"context"
	"tobialbertino/portfolio-be/internal/to_do/models/entity"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/mock"
)

type ToDoRepositoryMock struct {
	mock.Mock
}

func (m *ToDoRepositoryMock) Create(ctx context.Context, db *pgxpool.Pool, toDo *entity.ToDo) (int64, error) {
	args := m.Called(toDo)
	return 1, args.Error(1)
}

func (m *ToDoRepositoryMock) Update(ctx context.Context, db *pgxpool.Pool, toDo *entity.ToDo) (int64, error) {
	args := m.Called(toDo)
	return 1, args.Error(1)
}

func (m *ToDoRepositoryMock) Delete(ctx context.Context, db *pgxpool.Pool, id *int64) (int64, error) {
	args := m.Called(id)
	return 1, args.Error(1)
}

func (m *ToDoRepositoryMock) DeleteAll(ctx context.Context, db *pgxpool.Pool) (int64, error) {
	args := m.Called()
	return 1, args.Error(1)
}

func (m *ToDoRepositoryMock) GetAll(ctx context.Context, db *pgxpool.Pool) (*entity.ListToDo, error) {
	args := m.Called()
	var res entity.ListToDo = entity.ListToDo{{
		Id:         0,
		Title:      "",
		Status:     false,
		Created_at: 0,
		Updated_at: 0,
	}}
	return &res, args.Error(1)
}
