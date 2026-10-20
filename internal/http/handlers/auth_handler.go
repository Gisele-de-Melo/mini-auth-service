package handlers

import (
	"encoding/json"
	"net/http"

	"mini-auth-service/internal/service"
	"mini-auth-service/pkg/logger"
)

type AuthHandler struct {
	authService *service.AuthService
	log         *logger.Logger
}

func NewAuthHandler(authService *service.AuthService, log *logger.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		log:         log,
	}
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email e senha são obrigatórios", http.StatusBadRequest)
		return
	}

	err := h.authService.Register(req.Email, req.Password)
	if err != nil {
		if err == service.ErrEmailAlreadyUsed {
			http.Error(w, "Email já está em uso", http.StatusConflict)
			return
		}
		h.log.Error("Erro ao registrar usuário", "error", err)
		http.Error(w, "Erro interno ao registrar", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"Usuário registrado com sucesso"}`))
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		if err == service.ErrInvalidCredentials {
			http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
			return
		}
		h.log.Error("Erro ao fazer login", "error", err)
		http.Error(w, "Erro interno ao autenticar", http.StatusInternalServerError)
		return
	}

	resp := map[string]string{
		"token": token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
