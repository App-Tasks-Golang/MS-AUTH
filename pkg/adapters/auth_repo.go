package adapters

import (
	"OAuth-Service-Go/pkg/domain"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *domain.Auth) error
	FindUserByUsername(username string) (*domain.Auth, error)
	ExistsUser(email string) (bool, error)
	AddTokenToBlacklist(token string) error
	IsTokenBlacklisted(token string) (bool, error)
}

// Implementación de AuthRepository
type AuthRepo struct {
	db *gorm.DB
}

// Nueva instancia del repositorio
func NewAuthRepository(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

// Crear un nuevo usuario
func (r *AuthRepo) CreateUser(user *domain.Auth) error {
	return r.db.Create(user).Error
}

// Buscar usuario por email
func (r *AuthRepo) FindUserByUsername(email string) (*domain.Auth, error) {
	var user domain.Auth
	err := r.db.Where("username = ?", email).First(&user).Error
	return &user, err
}



func (r *AuthRepo) ExistsUser(email string) (bool, error) {
	var count int64
	// Contar los registros que tienen el email
	err := r.db.Model(&domain.Auth{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}




// Agregar un token a la lista negra
func (r *AuthRepo) AddTokenToBlacklist(token string) error {
	blacklistToken := &domain.TokenBlacklist{Token: token}
	return r.db.Create(blacklistToken).Error
}

// Verificar si el token ya está en la lista negra
func (r *AuthRepo) IsTokenBlacklisted(token string) (bool, error) {
	var count int64
	err := r.db.Model(&domain.TokenBlacklist{}).Where("token = ?", token).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

