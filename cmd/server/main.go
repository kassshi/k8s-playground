package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"connectrpc.com/grpcreflect"
	"github.com/kassshi/golang-practice/gen/todo/v1/todov1connect"
	"github.com/kassshi/golang-practice/internal/config"
	"github.com/kassshi/golang-practice/internal/handler"
	"github.com/kassshi/golang-practice/internal/infra/db"
	"github.com/kassshi/golang-practice/internal/service"
)

func main() {
	mux := http.NewServeMux()
	todoHandler := handler.NewTodoHandler(service.NewTodoService())
	path, h := todov1connect.NewTodoServiceHandler(todoHandler)
	mux.Handle(path, h)

	// Reflection
	ref := grpcreflect.NewStaticReflector("todo.v1.TodoService")
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

	// database
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := db.NewPool(ctx, config.DataBaseUrl())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
