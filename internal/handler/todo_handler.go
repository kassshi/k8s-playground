package handler

import (
	"context"

	"buf.build/go/protovalidate"
	"connectrpc.com/connect"
	"github.com/google/uuid"
	v1 "github.com/kassshi/golang-practice/internal/gen/todo/v1"
	"github.com/kassshi/golang-practice/internal/gen/todo/v1/todov1connect"
	"github.com/kassshi/golang-practice/internal/infra/db/sqlc"
	"github.com/kassshi/golang-practice/internal/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) todov1connect.TodoServiceHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) CreateTodo(ctx context.Context, req *connect.Request[v1.CreateTodoRequest]) (*connect.Response[v1.Todo], error) {
	if err := protovalidate.Validate(req.Msg); err != nil {
		return nil, err
	}
	result, err := h.service.CreateTodo(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(toProtoTodo(result)), nil
}
func (h *TodoHandler) GetTodo(ctx context.Context, req *connect.Request[v1.GetTodoRequest]) (*connect.Response[v1.Todo], error) {
	if err := protovalidate.Validate(req.Msg); err != nil {
		return nil, err
	}

	result, err := h.service.GetTodoByID(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(toProtoTodo(result)), nil
}

func (h *TodoHandler) UpdateTodo(ctx context.Context, req *connect.Request[v1.UpdateTodoRequest]) (*connect.Response[v1.Todo], error) {
	if err := protovalidate.Validate(req.Msg); err != nil {
		return nil, err
	}

	if err := h.service.UpdateTodoStatus(ctx, req); err != nil {
		return nil, err
	}
	return connect.NewResponse(req.Msg.GetTodo()), nil
}

func (h *TodoHandler) DeleteTodo(ctx context.Context, req *connect.Request[v1.DeleteTodoRequest]) (*connect.Response[emptypb.Empty], error) {
	if err := protovalidate.Validate(req.Msg); err != nil {
		return nil, err
	}

	if err := h.service.DeleteTodo(ctx, req); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
func (h *TodoHandler) ListTodos(ctx context.Context, req *connect.Request[v1.ListTodosRequest]) (*connect.Response[v1.ListTodosResponse], error) {
	if err := protovalidate.Validate(req.Msg); err != nil {
		return nil, err
	}

	result, err := h.service.ListTodosByUserID(ctx)
	if err != nil {
		return nil, err
	}
	var todos []*v1.Todo
	for _, todo := range result {
		todos = append(todos, toProtoTodo(todo))
	}
	return connect.NewResponse(&v1.ListTodosResponse{Todos: todos}), nil
}

func toProtoTodo(t sqlc.Todo) *v1.Todo {
	return &v1.Todo{
		Name:        uuid.UUID(t.ID.Bytes).String(),
		Title:       t.Title,
		Description: t.Description,
		Status:      v1.Status(v1.Status_value["STATUS_"+t.Status]),
	}
}
