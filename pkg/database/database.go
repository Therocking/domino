package database

import (
	"fmt"
	"githup/Therocking/dominoes/internal/entities"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewConfig() *Config {
	return &Config{
		Host:     getEnv("DB_HOST"),
		Port:     getEnv("DB_PORT"),
		User:     getEnv("DB_USER"),
		Password: getEnv("DB_PASSWORD"),
		DBName:   getEnv("DB_NAME"),
		SSLMode:  getEnv("DB_SSLMODE"),
	}
}

// Connect establece la conexión con la base de datos
func Connect() (*gorm.DB, error) {
	config := NewConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	// Configurar logger para GORM
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Umbral para queries lentos
			LogLevel:                  logger.Info, // Nivel de log
			IgnoreRecordNotFoundError: true,        // Ignorar errores de "record not found"
			Colorful:                  true,        // Habilitar colores
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(
		&entities.Session{},
		&entities.Team{},
		&entities.Game{},
		&entities.GamePoint{},
		&entities.Ranking{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	// Configurar conexión pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Configuración del pool de conexiones
	sqlDB.SetMaxIdleConns(10)           // Conexiones inactivas máximas
	sqlDB.SetMaxOpenConns(100)          // Conexiones abiertas máximas
	sqlDB.SetConnMaxLifetime(time.Hour) // Tiempo máximo de vida de una conexión

	log.Println("Successfully connected to database!")
	return db, nil
}

// getEnv obtiene una variable de entorno o devuelve un valor por defecto
func getEnv(key string) string {
	value := os.Getenv(key)

	return value
}
