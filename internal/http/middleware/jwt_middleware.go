package middleware

import (
	"context"
	"net/http"
	"strings"

	"mini-auth-service/internal/security"
)

type contextKey string

const (
	ContextUserIDKey contextKey = "user_id"
	ContextEmailKey  contextKey = "email"
)

func JWTMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header ausente", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Formato do token inválido", http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]

			// Validar token
			claims, err := security.ValidateJWT(tokenString, secret)
			if err != nil {
				http.Error(w, "Token inválido ou expirado", http.StatusUnauthorized)
				return
			}

			// Extrair claims
			userID := int(claims["user_id"].(float64))
			email := claims["email"].(string)

			// Colocar no contexto
			ctx := context.WithValue(r.Context(), ContextUserIDKey, userID)
			ctx = context.WithValue(ctx, ContextEmailKey, email)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
