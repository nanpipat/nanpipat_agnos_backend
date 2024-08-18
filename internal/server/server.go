package server

import (
	"fmt"
	"time"

	"github.com/nanpipat/nanpipat-agnos-backend/internal/config"
	"github.com/nanpipat/nanpipat-agnos-backend/internal/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	config *config.Config
}

func New(cfg *config.Config) *Server {
	db, err := initDBWithRetry(cfg)
	if err != nil {
		panic(err)
	}

	// Run migrations
	if err := migrations.Migrate(db); err != nil {
		panic(fmt.Sprintf("Failed to run migrations: %v", err))
	}

	return &Server{
		db:     db,
		config: cfg,
	}
}

func (s *Server) DB() *gorm.DB {
	return s.db
}

func initDBWithRetry(cfg *config.Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err = initDB(cfg)
		if err == nil {
			return db, nil
		}
		fmt.Printf("Failed to connect to database. Retrying in 5 seconds... (Attempt %d/%d)\n", i+1, maxRetries)
		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect to database after %d attempts", maxRetries)
}

func initDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
