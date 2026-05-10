package main

import (
	"net/http"

	"github.com/Yashk767/internal/handler"
	"github.com/Yashk767/internal/repository"
	"github.com/Yashk767/internal/service"
)

func main() {
	//setup repository
	repository := repository.NewInMemoryRepository()

	//setup service
	service := service.NewService(repository)

	//setup handler
	handler := handler.NewHandler(service)

	//setup server
	router := http.NewServeMux()

}
