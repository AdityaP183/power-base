package routes

import (
	"github.com/AdityaP183/power-base/internal/api/handlers"
	"github.com/AdityaP183/power-base/internal/api/middleware"
	"github.com/AdityaP183/power-base/internal/cache"

	"time"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(
	router *gin.Engine,
	heroHandler *handlers.HeroHandler,
	redisCache *cache.RedisCache,
) {
	api := router.Group("")

	api.Use(middleware.RateLimiterMiddleware(10, 20))

	heroes := api.Group("/heroes")
	{
		heroes.GET("",
			middleware.CacheMiddleware(redisCache, 5*time.Second),
			middleware.PaginationMiddleware(),
			heroHandler.GetHeroes,
		)
		heroes.GET("/:id", heroHandler.GetHeroByID)
	}

	// TODO: Docs api
	// docs := router.Group("/docs")
	// {
	// 	// Serve static documentation site
	// 	docs.Static("/", "./web/static")

	// 	// Route for API documentation
	// 	docs.GET("/heroes", func(c *gin.Context) {
	// 		c.File("./web/templates/docs/heroes.html")
	// 	})
	// }
}
