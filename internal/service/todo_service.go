// Package service fot todo
package service

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	v1 "github.com/kassshi/golang-practice/internal/gen/todo/v1"
	"github.com/kassshi/golang-practice/internal/infra/db/sqlc"
	"github.com/kassshi/golang-practice/internal/repository"
)

// TODO: 認証実装後にログインユーザーのIDに差し替える
var tempUserID = pgtype.UUID{Bytes: uuid.MustParse("00000000-0000-0000-0000-000000000001"), Valid: true}

type TodoService struct {
	repository *repository.TodoRepository
}

func NewTodoService(repository *repository.TodoRepository) *TodoService {
	return &TodoService{
		repository: repository,
	}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *connect.Request[v1.CreateTodoRequest]) (sqlc.Todo, error) {
	todo := req.Msg.GetTodo()
	id := uuid.New()
	args := sqlc.CreateTodoParams{
		ID:          pgtype.UUID{Bytes: id, Valid: true},
		UserID:      tempUserID,
		Title:       todo.GetTitle(),
		Description: todo.GetDescription(),
		Status:      toDBStatus(todo.GetStatus()),
	}
	return s.repository.CreateTodo(ctx, args)
}

func (s *TodoService) GetTodoByID(ctx context.Context, req *connect.Request[v1.GetTodoRequest]) (sqlc.Todo, error) {
	id, err := uuid.Parse(req.Msg.GetName())
	if err != nil {
		return sqlc.Todo{}, fmt.Errorf("invalid uuid: %w", err)
	}
	args := sqlc.GetTodoByIDParams{
		ID:     pgtype.UUID{Bytes: id, Valid: true},
		UserID: tempUserID,
	}
	return s.repository.GetTodoByID(ctx, args)
}

func (s *TodoService) UpdateTodoStatus(ctx context.Context, req *connect.Request[v1.UpdateTodoRequest]) error {
	todo := req.Msg.GetTodo()
	id, err := uuid.Parse(todo.GetName())
	if err != nil {
		return fmt.Errorf("invalid uuid: %w", err)
	}
	args := sqlc.UpdateTodoStatusParams{
		ID:          pgtype.UUID{Bytes: id, Valid: true},
		UserID:      tempUserID,
		Title:       todo.GetTitle(),
		Description: todo.GetDescription(),
		Status:      toDBStatus(todo.GetStatus()),
	}
	return s.repository.UpdateTodoStatus(ctx, args)
}

func (s *TodoService) DeleteTodo(ctx context.Context, req *connect.Request[v1.DeleteTodoRequest]) error {
	id, err := uuid.Parse(req.Msg.GetName())
	if err != nil {
		return fmt.Errorf("invalid uuid: %w", err)
	}
	args := sqlc.DeleteTodoParams{
		ID:     pgtype.UUID{Bytes: id, Valid: true},
		UserID: tempUserID,
	}
	return s.repository.DeleteTodo(ctx, args)
}

func (s *TodoService) ListTodosByUserID(ctx context.Context) ([]sqlc.Todo, error) {
	return s.repository.ListTodosByUserID(ctx, tempUserID)
}

func toDBStatus(s v1.Status) string {
	switch s {
	case v1.Status_STATUS_TODO:
		return "TODO"
	case v1.Status_STATUS_IN_PROGRESS:
		return "IN_PROGRESS"
	case v1.Status_STATUS_DONE:
		return "DONE"
	default:
		return "TODO"
	}
}

