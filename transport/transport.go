package transport

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authService *AuthHandler) {
	// Definir las rutas de usuarios
	auth := router.Group("/auth")
	{
		auth.POST("/register", authService.Register)
		auth.POST("/login", authService.Login)
		auth.POST("/logout", authService.Logout)
	}
}


