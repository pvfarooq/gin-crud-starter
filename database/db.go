package database

import (
	"fmt"

	"github.com/pvfarooq/go-gin-crud/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes and returns a new database connection using GORM.
// It takes a configuration object that contains PostgreSQL connection details.
// The function constructs a connection string (DSN) using the provided config
// and attempts to establish a connection to the database.
//
// Parameters:
//   - cfg: A pointer to config.Config containing database connection parameters
//
// Returns:
//   - *gorm.DB: A pointer to the initialized database connection
//
// Panics:
//   - If the database connection fails to establish
func InitDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.POSTGRES_HOST, cfg.POSTGRES_USER, cfg.POSTGRES_PASSWORD, cfg.POSTGRES_DB, cfg.POSTGRES_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Uncomment this line to auto migrate the registered models
	// db.AutoMigrate(&models.User{})
	return db
}
