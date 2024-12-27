package auth

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("tok3n2024") // Clave privada para firmar tokens.

// Estructura para los claims del token
type Claims struct {
    UserID uint `json:"user_id"`
    Username string `json:"username"`
    jwt.RegisteredClaims
}

// Función para generar un token
func GenerateToken(userID uint, username string) (string, error) {
    claims := &Claims{
        Username: username,
        UserID: userID,

        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    fmt.Printf("Estructura del token: %+v\n", claims)
    return token.SignedString(jwtKey)
}


// Función para validar un token
func ValidateToken(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil || !token.Valid {
        return nil, err
    }
    return claims, nil
}
