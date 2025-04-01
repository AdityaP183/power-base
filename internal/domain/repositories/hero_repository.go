package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/AdityaP183/power-base/internal/domain/models"
	"gorm.io/gorm"
)

// HeroRepository handles database operations for heroes
type HeroRepository struct {
	db *gorm.DB
}

// NewHeroRepository creates a new hero repository
func NewHeroRepository(db *gorm.DB) *HeroRepository {
	return &HeroRepository{db: db}
}

// GetAll retrieves heroes with pagination
func (r *HeroRepository) GetAll(ctx context.Context, page, limit int) ([]models.Hero, int64, error) {
	var heroes []models.Hero
	var total int64

	if err := r.db.Model(&models.Hero{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculating offset
	offset := (page - 1) * limit

	result := r.db.Offset(offset).Limit(limit).Find(&heroes)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return heroes, total, nil
}

// GetByID retrieves a hero by ID
func (r *HeroRepository) GetByID(ctx context.Context, id string, fields []string) (*models.Hero, error) {
	var hero models.Hero

	query := r.db

	// If specific fields are requested, select only those
	if len(fields) > 0 {
		fields = append(fields, "id")

		// Convert fields to database column names if needed
		var selectedFields []string
		for _, field := range fields {
			switch field {
			case "powerstats":
				selectedFields = append(selectedFields, "powerstats")
			case "biography":
				selectedFields = append(selectedFields, "biography")
			case "appearance":
				selectedFields = append(selectedFields, "appearance")
			case "connections":
				selectedFields = append(selectedFields, "connections")
			case "work":
				selectedFields = append(selectedFields, "work")
			case "image":
				selectedFields = append(selectedFields, "image")
			}
		}

		if len(selectedFields) > 0 {
			query = query.Select(selectedFields)
		}
	}

	if err := query.Where("id = ?", id).First(&hero).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("hero not found: %v", id)
		}
	}

	return &hero, nil
}

// CreateHero adds a new hero
func (r *HeroRepository) CreateHero(ctx context.Context, hero *models.Hero) error {
	return r.db.Create(hero).Error
}

// UpdateHero adds a new hero
func (r *HeroRepository) UpdateHero(ctx context.Context, hero *models.Hero) error {
	return r.db.Save(hero).Error
}

// DeleteHero adds a new hero
func (r *HeroRepository) DeleteHero(ctx context.Context, id string) error {
	return r.db.Delete(&models.Hero{}, "id = ?", id).Error
}
