package main

import (
	"fmt"
	"net/http"

	"github.com/hasanarofid/golang-backend/internal/config"
	"github.com/hasanarofid/golang-backend/internal/handler"
	"github.com/hasanarofid/golang-backend/internal/repository"
	"github.com/hasanarofid/golang-backend/internal/service"
)

func main() {
	db := config.ConnectDB()

	userRepo := &repository.UserRepository{DB: db}
	userService := &service.UserService{Repo: userRepo}
	userHandler := &handler.UserHandler{Service: userService}

	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/login", userHandler.Login)

	fmt.Println("server running on :8080")
	http.ListenAndServe(":8080", nil)
}
