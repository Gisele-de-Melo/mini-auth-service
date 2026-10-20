package handlers

import (
	"encoding/json"
	"net/http"

	"mini-auth-service/internal/http/middleware"
	"mini-auth-service/internal/service"
	"mini-auth-service/pkg/logger"
)

type UserHandler struct {
	userService *service.UserService
	log         *logger.Logger
}

func NewUserHandler(userService *service.UserService, log *logger.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		log:         log,
	}
}

// GET /me
func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.ContextUserIDKey)
	email := r.Context().Value(middleware.ContextEmailKey)

	if userID == nil || email == nil {
		http.Error(w, "Usuário não autenticado", http.StatusUnauthorized)
		return
	}

	user, err := h.userService.GetByID(userID.(int))
	if err != nil || user == nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	resp := map[string]interface{}{
		"id":         user.ID,
		"email":      user.Email,
		"created_at": user.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
