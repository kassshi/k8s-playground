package service

import (
	"context"
	"errors"
	"time"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
	v1 "github.com/kassshi/golang-practice/internal/gen/auth/v1"
	"github.com/kassshi/golang-practice/internal/infra/db/sqlc"
	"github.com/kassshi/golang-practice/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repository.UserRepository
	jwtSecret      string
}

func NewAuthService(userRepository *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		jwtSecret:      jwtSecret,
	}
}

func (s *AuthService) Signup(ctx context.Context, req *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error) {
	if err := s.validateSignupRequest(req.Msg); err != nil {
		return nil, err
	}

	passwordHash, err := s.hashPassword(req.Msg.GetPassword())
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.Create(ctx, sqlc.CreateUserParams{
		Email:        req.Msg.Email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.SignupResponse{AccessToken: s.generateAccessToken(user)}), nil
}

func (s *AuthService) validateSignupRequest(req *v1.SignupRequest) error {
	if req.GetPassword() != req.GetConfirmPassword() {
		return errors.New("password and confirm password do not match")
	}
	return nil
}

func (s *AuthService) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *AuthService) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	user, err := s.userRepository.GetUserByEmail(ctx, req.Msg.GetEmail())
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Msg.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return connect.NewResponse(&v1.LoginResponse{AccessToken: s.generateAccessToken(user)}), err
}

func (s *AuthService) generateAccessToken(user sqlc.User) string {
	claim := jwt.RegisteredClaims{
		Subject:   user.ID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	accessToken, _ := token.SignedString([]byte(s.jwtSecret))
	return accessToken
}
