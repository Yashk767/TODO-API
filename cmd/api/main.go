package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Yashk767/internal/handler"
	"github.com/Yashk767/internal/repository"
	"github.com/Yashk767/internal/server"
	"github.com/Yashk767/internal/service"
)

func main() {
	//setup repository

	//repository := repository.NewInMemoryRepository()
	repository, err := repository.NewSqliteRepository()
	if err != nil {
		log.Fatal("Error setting up repository: ", err)
	}

	//setup service
	service := service.NewService(repository)

	//setup handler
	handler := handler.NewHandler(service)

	//setup router
	router := server.NewRouter(handler)

	httpServer := &http.Server{
		Addr:         "localhost:8082",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("server started on :", httpServer.Addr)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
