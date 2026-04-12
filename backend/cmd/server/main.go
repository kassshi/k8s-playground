package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/validate"
	"github.com/kassshi/golang-practice/backend/internal/config"
	"github.com/kassshi/golang-practice/backend/internal/gen/auth/v1/authv1connect"
	"github.com/kassshi/golang-practice/backend/internal/gen/todo/v1/todov1connect"
	"github.com/kassshi/golang-practice/backend/internal/gen/user/v1/userv1connect"
	"github.com/kassshi/golang-practice/backend/internal/handler"
	"github.com/kassshi/golang-practice/backend/internal/infra/db"
	"github.com/kassshi/golang-practice/backend/internal/infra/db/sqlc"
	"github.com/kassshi/golang-practice/backend/internal/middleware"
	"github.com/kassshi/golang-practice/backend/internal/repository"
	"github.com/kassshi/golang-practice/backend/internal/service"
)

func main() {

	// Config
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := db.NewPool(ctx, config.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// repository
	queries := sqlc.New(pool)
	userRepository := repository.NewUserRepository(queries)
	todoRepository := repository.NewTodoRepository(queries)

	// service
	todoService := service.NewTodoService(todoRepository)
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository, config.JwtSecret)

	// handler
	// Create the validation interceptor provided by connectrpc.com/validate.
	validateInterceptor := validate.NewInterceptor()
	mux := http.NewServeMux()
	todoHandler := handler.NewTodoHandler(todoService)
	userHandler := handler.NewUserHandler(userService)
	oauthHandler := handler.NewAuthHandler(authService)
	path, h := todov1connect.NewTodoServiceHandler(todoHandler, connect.WithInterceptors(validateInterceptor))
	mux.Handle(path, middleware.AuthMiddleware(config.JwtSecret)(h))
	path, h = userv1connect.NewUserServiceHandler(userHandler)
	mux.Handle(path, h)
	path, h = authv1connect.NewAuthServiceHandler(oauthHandler)
	mux.Handle(path, h)

	// Reflection
	ref := grpcreflect.NewStaticReflector("todo.v1.TodoService", "auth.v1.AuthService", "user.v1.UserService")
	mux.Handle(grpcreflect.NewHandlerV1(ref))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(ref))

	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true)
	s := http.Server{
		Addr:      ":8080",
		Handler:   mux,
		Protocols: p,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
