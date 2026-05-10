package server

import (
	"net/http"

	"github.com/Yashk767/internal/handler"
)

func NewRouter(todoHandler *handler.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /todos", todoHandler.CreateTodo)
	mux.HandleFunc("GET /todos/{id}", todoHandler.GetToDo)
	mux.HandleFunc("PUT /todos/{id}", todoHandler.UpdateTodo)
	mux.HandleFunc("DELETE /todos/{id}", todoHandler.DeleteTodo)
	mux.HandleFunc("GET /todos", todoHandler.GetAllToDos)

	return mux
}
