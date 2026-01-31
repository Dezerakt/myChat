package pg

import (
	"fmt"
	"log"
	"myChat/config"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func NewConnection(cfg *config.Postgres) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return nil, err
	}

	log.Println("Connected to PostgreSQL successfully")
	return db, nil
}
