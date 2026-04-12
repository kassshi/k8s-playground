// Package service fot todo
package service

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	v1 "github.com/kassshi/golang-practice/backend/internal/gen/todo/v1"
	"github.com/kassshi/golang-practice/backend/internal/infra/db/sqlc"
	"github.com/kassshi/golang-practice/backend/internal/middleware"
	"github.com/kassshi/golang-practice/backend/internal/repository"
)

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
	userID, err := getUserIDFromcontext(ctx)
	if err != nil {
		return sqlc.Todo{}, fmt.Errorf("failed to get user id from context: %w", err)
	}
	args := sqlc.CreateTodoParams{
		ID:          pgtype.UUID{Bytes: id, Valid: true},
		UserID:      userID,
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
	userID, err := getUserIDFromcontext(ctx)
	if err != nil {
		return sqlc.Todo{}, err
	}
	args := sqlc.GetTodoByIDParams{
		ID:     pgtype.UUID{Bytes: id, Valid: true},
		UserID: userID,
	}
	return s.repository.GetTodoByID(ctx, args)
}

func (s *TodoService) UpdateTodoStatus(ctx context.Context, req *connect.Request[v1.UpdateTodoRequest]) error {
	todo := req.Msg.GetTodo()
	id, err := uuid.Parse(todo.GetName())
	if err != nil {
		return fmt.Errorf("invalid uuid: %w", err)
	}
	userID, err := getUserIDFromcontext(ctx)
	if err != nil {
		return err
	}
	args := sqlc.UpdateTodoStatusParams{
		ID:          pgtype.UUID{Bytes: id, Valid: true},
		UserID:      userID,
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
	userID, err := getUserIDFromcontext(ctx)
	if err != nil {
		return err
	}
	args := sqlc.DeleteTodoParams{
		ID:     pgtype.UUID{Bytes: id, Valid: true},
		UserID: userID,
	}
	return s.repository.DeleteTodo(ctx, args)
}

func (s *TodoService) ListTodosByUserID(ctx context.Context) ([]sqlc.Todo, error) {
	userID, err := getUserIDFromcontext(ctx)
	if err != nil {
		return nil, err
	}
	return s.repository.ListTodosByUserID(ctx, userID)
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

func getUserIDFromcontext(ctx context.Context) (pgtype.UUID, error) {
	userIDStr := ctx.Value(middleware.UserIDKey).(string)
	id, err := uuid.Parse(userIDStr)
	if err != nil {
		return pgtype.UUID{}, fmt.Errorf("invalid uuid: %w", err)
	}
	return pgtype.UUID{Bytes: id, Valid: true}, nil
}
