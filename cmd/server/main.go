package main

import (
	"log"
	"net/http"

	"connectrpc.com/grpcreflect"
	"github.com/kassshi/golang-practice/gen/todo/v1/todov1connect"
	"github.com/kassshi/golang-practice/internal/handler"
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
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
