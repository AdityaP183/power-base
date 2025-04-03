package services

import (
	"context"

	"github.com/AdityaP183/power-base/internal/domain/models"
	"github.com/AdityaP183/power-base/internal/domain/repositories"
)

// HeroService handles business logic for heroes
type HeroService struct {
	repo  *repositories.HeroRepository
	cache Cache
}

type Cache interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, ttl int) error
	Delete(key string) error
}

// NewHeroService creates a new hero service
func NewHeroService(repo *repositories.HeroRepository, cache Cache) *HeroService {
	return &HeroService{
		repo:  repo,
		cache: cache,
	}
}

// GetHeroes retrieves all heroes with pagination
func (s *HeroService) GetHeroes(ctx context.Context, page, limit int) ([]models.Hero, int64, error) {
	return s.repo.GetAll(ctx, page, limit)
}

// GetHeroByID retrieves a hero by ID
func (s *HeroService) GetHeroByID(ctx context.Context, id string, fields []string) (*models.Hero, error) {
	return s.repo.GetByID(ctx, id, fields)
}

// CreateHero adds a new hero
func (s *HeroService) CreateHero(ctx context.Context, hero *models.Hero) error {
	return s.repo.CreateHero(ctx, hero)
}

// UpdateHero updates an existing hero
func (s *HeroService) UpdateHero(ctx context.Context, hero *models.Hero) error {
	err := s.repo.UpdateHero(ctx, hero)

	if err != nil {
		_ = s.cache.Delete("hero" + hero.ID)
		_ = s.cache.Delete("hero:list")
	}
	return err
}

// DeleteHero removes a hero
func (s *HeroService) DeleteHero(ctx context.Context, id string) error {
	err := s.repo.DeleteHero(ctx, id)

	if err != nil {
		_ = s.cache.Delete("hero" + id)
		_ = s.cache.Delete("hero:list")
	}
	return err
}
