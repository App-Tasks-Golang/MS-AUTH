package config

import (
	"OAuth-Service-Go/pkg/domain"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	// Cargar variables del archivo .env
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error cargando el archivo .env: %w", err)
	}

	// Formar la cadena de conexión a la base de datos para MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("AUTH_ROOT"),     // Usuario de la base de datos
		os.Getenv("AUTH_PASSWORD"), // Contraseña de la base de datos
		os.Getenv("AUTH_HOST"),     // Dirección del host de la base de datos
		os.Getenv("AUTH_PORT"),     // Puerto de la base de datos
		os.Getenv("AUTH_NAME"),     // Nombre de la base de datos
	)

	// Intentar abrir la conexión a la base de datos
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Cambiar a MySQL
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %w", err)
	}


	if err := db.AutoMigrate(&domain.Auth{}, &domain.TokenBlacklist{}); err != nil {
		return nil, fmt.Errorf("error al realizar la migración: %w", err)
	}

	// Devolver la conexión exitosa
	return db, nil
}
