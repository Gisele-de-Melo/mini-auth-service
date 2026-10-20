package service

import (
	"errors"
	"time"

	"mini-auth-service/internal/model"
	"mini-auth-service/internal/repository"
	"mini-auth-service/internal/security"
)

var (
	ErrInvalidCredentials = errors.New("email ou senha inválidos")
	ErrEmailAlreadyUsed   = errors.New("email já está em uso")
)

type AuthService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// Registrar novo usuário
func (s *AuthService) Register(email, password string) error {
	// Verificar se email já existe
	existing, _ := s.userRepo.FindByEmail(email)
	if existing != nil {
		return ErrEmailAlreadyUsed
	}

	// Hash da senha
	hashed, err := security.HashPassword(password)
	if err != nil {
		return err
	}

	user := &model.User{
		Email:     email,
		Password:  hashed,
		CreatedAt: time.Now(),
	}

	return s.userRepo.Create(user)
}

// Login
func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil || user == nil {
		return "", ErrInvalidCredentials
	}

	// Verificar senha
	if !security.CheckPassword(password, user.Password) {
		return "", ErrInvalidCredentials
	}

	// Gerar JWT
	token, err := security.GenerateJWT(user.ID, user.Email, s.jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) JwtSecret() string {
	return s.jwtSecret
}
