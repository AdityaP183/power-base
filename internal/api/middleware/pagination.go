package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PaginationMiddleware handles pagination parameters
func PaginationMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil{
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Invalid page parameter",
			})
			c.Abort()
			return
		}

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
		if err != nil || limit < 1 || limit > 100{
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Invalid limit parameter",
			})
			c.Abort()
			return
		}

		c.Set("page", page)
		c.Set("limit", limit)

		c.Next()
	}
}
