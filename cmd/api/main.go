package main

import (
	"log"
	net "net/http"

	"mini-auth-service/config"
	"mini-auth-service/internal/http"
	"mini-auth-service/internal/repository"
	"mini-auth-service/internal/service"
	"mini-auth-service/pkg/logger"
)

func main() {
	// Carregar configuração
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	// Criar logger
	logg := logger.New()

	// Conectar ao banco SQLite
	db, err := repository.NewSQLiteDB(cfg.DatabasePath)
	if err != nil {
		logg.Fatal("erro ao conectar ao banco", "error", err)
	}

	// Criar repositórios
	userRepo := repository.NewUserRepository(db)

	// Criar serviços
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	userService := service.NewUserService(userRepo)

	// Criar router
	router := http.NewRouter(authService, userService, logg)

	// Iniciar servidor
	logg.Info("Servidor iniciado", "port", cfg.Port)

	if err := net.ListenAndServe(":"+cfg.Port, router); err != nil {
		logg.Fatal("erro ao iniciar servidor", "error", err)
	}
}
