package handler

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/kassshi/golang-practice/internal/gen/auth/v1"
	"github.com/kassshi/golang-practice/internal/gen/auth/v1/authv1connect"
	"github.com/kassshi/golang-practice/internal/service"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) authv1connect.AuthServiceHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Signup(ctx context.Context, req *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error) {
	return h.service.Signup(ctx, req)
}
func (h *AuthHandler) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return h.service.Login(ctx, req)
}
