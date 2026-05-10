package main

import (
	"net/http"
	"time"

	"github.com/Yashk767/internal/handler"
	"github.com/Yashk767/internal/repository"
	"github.com/Yashk767/internal/server"
	"github.com/Yashk767/internal/service"
)

func main() {
	//setup repository
	repository := repository.NewInMemoryRepository()

	//setup service
	service := service.NewService(repository)

	//setup handler
	handler := handler.NewHandler(service)

	//setup router
	router := server.NewRouter(handler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	server.ListenAndServe()
}
