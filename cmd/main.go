package main

import (
	"OAuth-Service-Go/internal/config"
	"OAuth-Service-Go/pkg/adapters"
	"OAuth-Service-Go/pkg/service"
	"OAuth-Service-Go/transport"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectToDB()
	if err != nil {
		panic("no se pudo conectar a la base de datos")
	}

	authRepo := adapters.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	handler := transport.NewAuthHandler(authService)

	r := gin.Default()
	transport.SetupRoutes(r, handler)

	r.Run(":8084")
}
