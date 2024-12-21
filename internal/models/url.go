package models

import "time"

// URL represents the structure of a shortened URL
type URL struct {
	ID        uint      `gorm:"primaryKey"`
	LongURL   string    `gorm:"not null"`
	ShortURL  string    `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Expiry    string    `gorm:"type:timestamp;"`
	UserID    uint      `gorm:"index"`
	Clicks    int       `gorm:"default:0"`
}
