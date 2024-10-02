package db

import (
	"fmt"
	"log"

	"github.com/parampreetr/grpc_go/config"
	"github.com/parampreetr/grpc_go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	env := config.GetEnvConfig()
	log.Println("Initializing Database")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Kolkata", env.PostgresHost, env.PostgresUser, env.PostgresPass, env.PostgresDB, env.PostgresPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Database initialized successfully")
	mitigate(db)

	return db, nil
}

func mitigate(db *gorm.DB) {
	log.Println("Mitigating Database...")
	db.AutoMigrate(&models.Task{})
}
