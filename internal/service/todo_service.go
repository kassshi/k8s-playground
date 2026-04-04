// Package service fot todo
package service

import (
	"context"

	"github.com/kassshi/golang-practice/internal/repository"
)

type TodoService struct {
	repository *repository.TodoRepository
}

func NewTodoService(repository *repository.TodoRepository) *TodoService {
	return &TodoService{
		repository: repository,
	}
}

func (s *TodoService) Ping(ctx context.Context) (string, error) {
	return "Pong", nil
}
