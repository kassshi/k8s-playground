// Package service fot todo
package service

import "context"

type TodoService struct{}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) Ping(ctx context.Context) (string, error) {
	return "Pong", nil
}
