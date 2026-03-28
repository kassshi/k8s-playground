package handler

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/kassshi/golang-practice/gen/todo/v1"
	"github.com/kassshi/golang-practice/gen/todo/v1/todov1connect"
	"github.com/kassshi/golang-practice/internal/service"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) todov1connect.TodoServiceHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) Ping(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	message, err := h.service.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&v1.PingResponse{Message: message}), nil
}
