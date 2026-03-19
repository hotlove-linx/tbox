package main

import (
	"fmt"
	"log"

	"tbox-backend/config"
	"tbox-backend/internal/model"
	"tbox-backend/internal/router"
	"tbox-backend/pkg/database"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize MySQL
	if err := database.InitMySQL(&cfg.Database); err != nil {
		log.Fatalf("Failed to initialize MySQL: %v", err)
	}

	// Auto migrate
	db := database.GetDB()
	if err := db.AutoMigrate(
		&model.User{},
		&model.Software{},
		&model.Category{},
		&model.Tag{},
		&model.SoftwareTag{},
		&model.SoftwareScreenshot{},
		&model.SoftwareVersion{},
		&model.SoftwareReview{},
		&model.UserSpace{},
	); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	// Initialize Redis
	if err := database.InitRedis(&cfg.Redis); err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}

	// Setup router
	r := router.SetupRouter(cfg.Server.Mode)

	// Start server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
