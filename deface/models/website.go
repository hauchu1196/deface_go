package models

import (
	"gorm.io/gorm"
	"time"
)

type Website struct {
	ID         uint      `gorm:"primaryKey"`
	URL        string    `gorm:"unique;not null"`
	HTMLHash   string    `gorm:"not null"`
	Screenshot string    `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Website{})
}

func GetAllWebsites(db *gorm.DB) ([]Website, error) {
	var websites []Website
	result := db.Find(&websites)
	return websites, result.Error
}

func (w *Website) Save(db *gorm.DB) error {
	return db.Save(w).Error
}
