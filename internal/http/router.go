package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"mini-auth-service/internal/http/handlers"
	jwtmw "mini-auth-service/internal/http/middleware"
	"mini-auth-service/internal/service"
	"mini-auth-service/pkg/logger"
)

func NewRouter(authService *service.AuthService, userService *service.UserService, log *logger.Logger) http.Handler {
	r := chi.NewRouter()

	// Middlewares globais
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService, log)
	userHandler := handlers.NewUserHandler(userService, log)

	// Rotas públicas
	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)

	// ROTAS PÚBLICAS (com rate limit)
	r.Group(func(public chi.Router) {
		public.Use(jwtmw.RateLimit())
		public.Post("/register", authHandler.Register)
		public.Post("/login", authHandler.Login)
	})

	// Rotas protegidas
	r.Group(func(protected chi.Router) {
		protected.Use(jwtmw.JWTMiddleware(authService.JwtSecret()))
		protected.Get("/me", userHandler.Me)
	})

	return r
}
