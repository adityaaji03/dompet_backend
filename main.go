package main

import (
	"log"

	"github.com/adityaaji03/dompet_backend/config"
	"github.com/adityaaji03/dompet_backend/database"
	"github.com/adityaaji03/dompet_backend/routes"
)

func main() {
	cfg := config.Load()

	db := database.InitMySQL(cfg)
	rdb := database.InitRedis(cfg)
	firebaseApp := database.InitFirebase(cfg)

	router := routes.Setup(db, rdb, firebaseApp, cfg)

	log.Printf("Server running on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
