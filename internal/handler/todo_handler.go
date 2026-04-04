package handler

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/kassshi/golang-practice/internal/gen/todo/v1"
	"github.com/kassshi/golang-practice/internal/gen/todo/v1/todov1connect"
	"github.com/kassshi/golang-practice/internal/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) todov1connect.TodoServiceHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) CreateTodo(context.Context, *connect.Request[v1.CreateTodoRequest]) (*connect.Response[v1.Todo], error) {
	return createResponse(), nil
}
func (h *TodoHandler) GetTodo(context.Context, *connect.Request[v1.GetTodoRequest]) (*connect.Response[v1.Todo], error) {
	return createResponse(), nil

}
func (h *TodoHandler) UpdateTodo(context.Context, *connect.Request[v1.UpdateTodoRequest]) (*connect.Response[v1.Todo], error) {
	return createResponse(), nil

}
func (h *TodoHandler) DeleteTodo(context.Context, *connect.Request[v1.DeleteTodoRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, nil

}
func (h *TodoHandler) ListTodos(context.Context, *connect.Request[v1.ListTodosRequest]) (*connect.Response[v1.ListTodosResponse], error) {
	return &connect.Response[v1.ListTodosResponse]{}, nil

}

func createResponse() *connect.Response[v1.Todo] {
	return connect.NewResponse(&v1.Todo{Name: "test"})
}
