package database

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB global variable
var DB *gorm.DB

// Config represents the database configuration
type Config struct {
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbPort     int    `env:"DB_PORT" envDefault:"5432"`
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"DB_NAME"`
}

// Connect initializes a connnection to the database thanks to a config object
func Connect(cfg Config) {
	var err error
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort)

	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to database!")
}
