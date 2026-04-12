package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kassshi/golang-practice/backend/internal/infra/db/sqlc"
)

type TodoRepository struct {
	querier sqlc.Querier
}

func NewTodoRepository(querier sqlc.Querier) *TodoRepository {
	return &TodoRepository{
		querier: querier,
	}
}

func (r *TodoRepository) CreateTodo(ctx context.Context, arg sqlc.CreateTodoParams) (sqlc.Todo, error) {
	return r.querier.CreateTodo(ctx, arg)
}
func (r *TodoRepository) DeleteTodo(ctx context.Context, arg sqlc.DeleteTodoParams) error {
	return r.querier.DeleteTodo(ctx, arg)
}
func (r *TodoRepository) GetTodoByID(ctx context.Context, arg sqlc.GetTodoByIDParams) (sqlc.Todo, error) {
	return r.querier.GetTodoByID(ctx, arg)
}
func (r *TodoRepository) ListTodosByUserID(ctx context.Context, userID pgtype.UUID) ([]sqlc.Todo, error) {
	return r.querier.ListTodosByUserID(ctx, userID)
}
func (r *TodoRepository) UpdateTodoStatus(ctx context.Context, arg sqlc.UpdateTodoStatusParams) error {
	return r.querier.UpdateTodoStatus(ctx, arg)
}
