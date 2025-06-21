package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"githup/Therocking/dominoes/internal/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

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

func Connect() error {
	config := NewConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	var err error

	DB, err = connectWithRetry(dsn, 10, newLogger)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = DB.AutoMigrate(
		&entities.Session{},
		&entities.Team{},
		&entities.Game{},
		&entities.GamePoint{},
		&entities.Ranking{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Successfully connected to database!")

	return nil
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func connectWithRetry(dsn string, maxRetries int, newLogger logger.Interface) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for i := range maxRetries {
		log.Printf("🔄 Intentando conexión a la base de datos (%d/%d)...", i+1, maxRetries)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})

		if err != nil {
			continue
		}

		sqlDB, err := db.DB()

		if err != nil {
			continue
		}

		err = sqlDB.Ping()
		if err == nil {
			return db, nil
		}

		log.Printf("❌ Fallo en conexión: %v", err)
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("❌ agotados los reintentos para conectar a la base de datos: %w", err)
}
