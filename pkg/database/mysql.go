package database

import (
	"time"

	"github.com/AdityaP183/power-base/internal/domain/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect establishes a connection to the MySQL database
func Connect(dbUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

// Migrate runs database migrations
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Hero{})
}
