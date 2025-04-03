package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// RateLimiterMiddleware implements rate limiting
func RateLimiterMiddleware(rps int, burst int) gin.HandlerFunc {
	rate := limiter.Rate{
		Period: time.Second,
		Limit:  int64(rps),
	}
	store := memory.NewStore()
	rateLimiter := limiter.New(store, rate)

	return func(c *gin.Context) {
		key := c.ClientIP()

		limiterCtx, err := rateLimiter.Get(c.Request.Context(), key)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Rate limiter error",
			})
			return
		}

		c.Header("X-RateLimit-Limit", strconv.FormatInt(limiterCtx.Limit, 10))
		c.Header("X-RateLimit-Remaining", strconv.FormatInt(limiterCtx.Remaining, 10))
		c.Header("X-RateLimit-Reset", strconv.FormatInt(limiterCtx.Reset, 10))

		if limiterCtx.Reached {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			return
		}

		c.Next()
	}
}
