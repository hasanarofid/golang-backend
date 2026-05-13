package handler

import (
	"encoding/json"
	"net/http"

	"golang-backend/internal/model"
	"golang-backend/internal/service"
)

type UserHandler struct {
	Service *service.UserService
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	err := h.Service.Register(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("register success"))
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input model.User
	json.NewDecoder(r.Body).Decode(&input)

	user, err := h.Service.Login(input.Email, input.Password)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	json.NewEncoder(w).Encode(user)
}
