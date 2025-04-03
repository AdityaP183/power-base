package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/AdityaP183/power-base/internal/cache"
	"github.com/gin-gonic/gin"
)

// responseBodyWriter is a custom response writer that captures the response body
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write captures the response and writes it to the underlying buffer
func (r *responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// CacheMiddleware handles response caching
func CacheMiddleware(redisCache *cache.RedisCache, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			c.Next()
			return
		}

		cacheKey := c.Request.URL.Path
		if c.Request.URL.RawQuery != "" {
			cacheKey += "?" + c.Request.URL.RawQuery
		}

		cachedResponse, err := redisCache.Get(cacheKey)
		if err != nil {
			var responseData map[string]any
			if err := json.Unmarshal(cachedResponse, &responseData); err == nil {
				c.Header("X-Cache", "HIT")
				c.JSON(http.StatusOK, responseData)
				c.Abort()
				return
			}
		}

		c.Header("X-Cache", "MISS")

		writer := &responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = writer

		c.Next()

		if c.Writer.Status() == http.StatusOK {
			if err := redisCache.Set(cacheKey, writer.body.Bytes(), int(duration.Seconds())); err != nil {
				c.JSON(http.StatusInternalServerError, err)
			}
		}
	}
}
