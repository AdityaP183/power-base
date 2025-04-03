package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/AdityaP183/power-base/internal/services"
	"github.com/AdityaP183/power-base/pkg/httputil"
	"github.com/gin-gonic/gin"
)

// HeroHandler handles HTTP requests for heroes
type HeroHandler struct {
	service *services.HeroService
}

// NewHeroHandler creates a new hero handler
func NewHeroHandler(service *services.HeroService) *HeroHandler {
	return &HeroHandler{service: service}
}

// GetHeroes handles GET /heroes
func (h *HeroHandler) GetHeroes(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	heroes, total, err := h.service.GetHeroes(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httputil.ErrorResponse{
			Error: "Failed to retrieve heroes",
		})
		return
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)
	nextPage := page + 1
	if int64(nextPage) > totalPages {
		nextPage = 0
	}

	c.JSON(http.StatusOK, gin.H{
		"heroes": heroes,
		"meta": gin.H{
			"total":      total,
			"page":       page,
			"limit":      limit,
			"totalPages": totalPages,
			"nextPage":   nextPage,
		},
	})
}

// GetHeroByID handles GET /heroes/:id
func (h *HeroHandler) GetHeroByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, httputil.ErrorResponse{Error: "hero ID is required"})
		return
	}

	fieldsParams := c.DefaultQuery("fields", "")
	var fields []string
	if fieldsParams != "" {
		fields = strings.Split(fieldsParams, ",")
	}

	hero, err := h.service.GetHeroByID(c.Request.Context(), id, fields)
	if err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			status = http.StatusNotFound
		}
		c.JSON(status, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hero": hero,
	})
}
