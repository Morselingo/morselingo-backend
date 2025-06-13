package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Morselingo/morselingo-backend/internal/model"
	"github.com/Morselingo/morselingo-backend/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var registerRequest model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := model.ValidateRegisterRequest(registerRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.service.RegisterUser(r.Context(), registerRequest); err != nil {
		if errors.Is(err, service.ErrorUserAlreadyExists) {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginRequest model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := model.ValidateLoginUserInput(loginRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := handler.service.LoginUser(r.Context(), loginRequest)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.LoginResponse{
		Token: token,
	})
}
