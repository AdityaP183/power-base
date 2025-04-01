package models

import (
	"time"
)

type Hero struct {
	ID          string      `json:"id" gorm:"primaryKey"`
	Name        string      `json:"name"`
	Appearance  Appearance  `json:"appearance" gorm:"embedded;embeddedPrefix:appearance_"`
	Biography   Biography   `json:"biography" gorm:"embedded;embeddedPrefix:biography_"`
	Connections Connections `json:"connections" gorm:"embedded;embeddedPrefix:connections_"`
	Powerstats  Powerstats  `json:"powerstats,omitempty" gorm:"embedded;embeddedPrefix:powerstats_"`
	Work        Work        `json:"work" gorm:"embedded;embeddedPrefix:work_"`
	Image       Image       `json:"image" gorm:"embedded;embeddedPrefix:image_"`
	Response    string      `json:"response"`
	CreatedAt   time.Time   `json:"-"`
	UpdatedAt   time.Time   `json:"-"`
}

type Appearance struct {
	EyeColor   string   `json:"eye-color" gorm:"column:appearance_eye_color"`
	Gender     string   `json:"gender"`
	HairColor  string   `json:"hair-color" gorm:"column:appearance_hair_color"`
	Height     []string `json:"height" gorm:"-"`
	HeightJSON string   `json:"-" gorm:"column:apperance_height"`
	Weight     []string `json:"weight" gorm:"-"`
	WeightJSON string   `json:"-" gorm:"column:apperance_weight"`
	Race       string   `json:"race"`
}

type Biography struct {
	Aliases         []string `json:"aliases" gorm:"-"`
	AliasesJSON     string   `json:"-" gorm:"column:biography_aliases"`
	Alignment       string   `json:"alignment"`
	AlterEgos       string   `json:"alter-egos" gorm:"column:biography_alter_egos"`
	FirstAppearance string   `json:"first-appearance" gorm:"column:biography_first_appearance"`
	FullName        string   `json:"full-name" gorm:"column:biography_full_name"`
	PlaceOfBirth    string   `json:"place-of-birth" gorm:"column:biography_place_of_birth"`
	Publisher       string   `json:"publisher"`
}

type Connections struct {
	GroupAffiliation string `json:"group-affiliation" gorm:"column:connections_group_affiliation"`
	Relatives        string `json:"relatives"`
}

type Powerstats struct {
	Combat       string `json:"combat"`
	Durability   string `json:"durability"`
	Intelligence string `json:"intelligence"`
	Power        string `json:"power"`
	Speed        string `json:"speed"`
	Strength     string `json:"strength"`
}

type Work struct {
	Base       string `json:"base"`
	Occupation string `json:"occupation"`
}

type Image struct {
	Url string `json:"url"`
}

// BeforeSave converts slice data to JSON for storage
func (h *Hero) BeforeSave() error {
	// Convert slices to JSON strings for storage
	// (Implementation would go here)
	return nil
}

// AfterFind converts JSON strings back to slices
func (h *Hero) AfterFind() error {
	// Convert JSON strings back to slices
	// (Implementation would go here)
	return nil
}
