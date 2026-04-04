package handler

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/kassshi/golang-practice/gen/user/v1"
	"github.com/kassshi/golang-practice/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return connect.NewResponse(&v1.PingResponse{Message: "Pong"}), nil
}
