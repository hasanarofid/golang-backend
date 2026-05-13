package main

import (
	"fmt"
	"net/http"

	"golang-backend/internal/config"
	"golang-backend/internal/handler"
	"golang-backend/internal/repository"
	"golang-backend/internal/service"
)

func main() {
	db := config.ConnectDB()

	userRepo := &repository.UserRepository{DB: db}
	userService := &service.UserService{Repo: userRepo}
	userHandler := &handler.UserHandler{Service: userService}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running 🚀"))
	})

	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/login", userHandler.Login)

	fmt.Println("server running on :8080")
	http.ListenAndServe(":8080", nil)
}
