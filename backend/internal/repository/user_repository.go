package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kassshi/golang-practice/backend/internal/infra/db/sqlc"
)

type UserRepository struct {
	querier sqlc.Querier
}

func NewUserRepository(querier sqlc.Querier) *UserRepository {
	return &UserRepository{
		querier: querier,
	}
}

func (r *UserRepository) Create(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error) {
	return r.querier.CreateUser(ctx, arg)
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (sqlc.User, error) {
	return r.querier.GetUserByEmail(ctx, email)
}

func (r *UserRepository) GetUserByID(ctx context.Context, id pgtype.UUID) (sqlc.User, error) {
	return r.querier.GetUserByID(ctx, id)
}
