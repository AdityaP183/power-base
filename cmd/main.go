package main

import (
	"log"

	"github.com/AdityaP183/power-base/config"
	"github.com/AdityaP183/power-base/internal/api/handlers"
	"github.com/AdityaP183/power-base/internal/api/routes"
	"github.com/AdityaP183/power-base/internal/cache"
	"github.com/AdityaP183/power-base/internal/domain/repositories"
	"github.com/AdityaP183/power-base/internal/services"
	"github.com/AdityaP183/power-base/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to run migration on the database: %s", err)
	}

	redisCache := cache.NewRedisCache(
		cfg.RedisAddr,
		cfg.RedisPass,
		cfg.RedisDB,
	)
	defer redisCache.Close()

	heroRepo := repositories.NewHeroRepository(db)
	heroService := services.NewHeroService(heroRepo, redisCache)
	heroHandler := handlers.NewHeroHandler(heroService)

	router := gin.Default()
	routes.SetupRoutes(router, heroHandler, redisCache)

	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
