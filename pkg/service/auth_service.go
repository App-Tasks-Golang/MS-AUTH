package service

import (
	"OAuth-Service-Go/auth"
	"OAuth-Service-Go/pkg/adapters"
	"OAuth-Service-Go/pkg/domain"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo adapters.AuthRepository
}

func NewAuthService(repo adapters.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

// Registro de usuario con hash
func (s *AuthService) RegisterUser(email, password, username string) (bool, string, error) {
	// Verificar si el usuario ya existe
	exist, err := s.repo.ExistsUser(email)
	if err != nil {
		// Error al consultar el repositorio
		return false, "", fmt.Errorf("error checking user existence: %w", err)
	}

	if exist {
		// Si el usuario ya existe
		return false, "email ya registrado anteriormente", nil
	}

	// Generar la contraseña cifrada con bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, "", fmt.Errorf("error hashing password: %w", err)
	}

	// Crear el objeto para el usuario
	user := &domain.Auth{
		Email:    email,
		Password: string(hashedPassword),
		Username: username,
	}

	// Guardar el usuario en el repositorio
	err = s.repo.CreateUser(user)
	if err != nil {
		return false, "", fmt.Errorf("error creating user: %w", err)
	}

	return true, "user successfully created", nil
}



// Validación de usuario (login)
func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.repo.FindUserByUsername(username)
	if err != nil {
		return "", errors.New("usuario no encontrado")
	}

	// Comparar contraseñas
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("credenciales incorrectas")
	}

	// Generar token
	token, err := auth.GenerateToken(uint(user.ID), user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}




// Validar un token de usuario
func (s *AuthService) ValidateToken(tokenString string) (*auth.Claims, error) {
	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		return nil, errors.New("token inválido o expirado")
	}
	return claims, nil
}



// Logout - Invalidar un token
func (s *AuthService) Logout(token string) error {
	// Aquí interactuamos con el repositorio para agregar el token a la lista negra
	err := s.repo.AddTokenToBlacklist(token)
	if err != nil {
		return fmt.Errorf("error al cerrar sesión: %w", err)
	}
	return nil
}