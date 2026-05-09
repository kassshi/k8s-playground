package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/validate"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	"github.com/rs/cors"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		execMigrate()
		return
	}
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
	cors := cors.New(cors.Options{
		AllowedOrigins: config.CorsAllowedOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Authorization", "Content-Type",
			"Connect-Protocol-Version"},
	})
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
		Handler:   cors.Handler(mux),
		Protocols: p,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func execMigrate() {

	// Config
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", config.DatabaseURL())

	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)

	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}
