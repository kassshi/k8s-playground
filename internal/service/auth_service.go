package service

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	v1 "github.com/kassshi/golang-practice/internal/gen/auth/v1"
	"github.com/kassshi/golang-practice/internal/infra/db/sqlc"
	"github.com/kassshi/golang-practice/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) Signup(ctx context.Context, req *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error) {
	if err := s.validateSignupRequest(req.Msg); err != nil {
		return nil, err
	}

	passwordHash, err := s.hashPassword(req.Msg)
	if err != nil {
		return nil, err
	}

	if _, err := s.userRepository.Create(ctx, sqlc.CreateUserParams{
		Email:        req.Msg.Email,
		PasswordHash: passwordHash,
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.SignupResponse{AccessToken: "test"}), nil
}

func (s *AuthService) validateSignupRequest(req *v1.SignupRequest) error {
	if req.GetPassword() != req.GetConfirmPassword() {
		return errors.New("password and confirm password do not match")
	}
	return nil
}

func (s *AuthService) hashPassword(req *v1.SignupRequest) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
