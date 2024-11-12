package models

import "time"

type Cheat struct {
	ID          int       `json:"id" gorm:"primaryKey"`        // Primary key
	Name        string    `json:"name" gorm:"unique;not null"` // Cheat name
	Content     string    `json:"content" gorm:"not null"`     // Cheat content
	Description string    `json:"description" gorm:"not null"` // Cheat description
	
	UpdatedAt   time.Time // When the cheat was last updated
	CreatedAt   time.Time // When the cheat was created
}
