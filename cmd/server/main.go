package main

import (
	"log"
	"net/http"

	"github.com/kassshi/golang-practice/gen/todo/v1/todov1connect"
	"github.com/kassshi/golang-practice/internal/handler"
	"github.com/kassshi/golang-practice/internal/service"
)

func main() {
	mux := http.NewServeMux()
	todoHandler := handler.NewTodoHandler(service.NewTodoService())
	path, h := todov1connect.NewTodoServiceHandler(todoHandler)
	mux.Handle(path, h)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
