package models

import "time"

type Key struct {
	ID         int       `json:"id" gorm:"primaryKey"`          // Primary key
	CheatID    int       `json:"cheat_id" gorm:"not null"`      // Cheat ID
	Key        string    `json:"key" gorm:"unique;not null"`    // Actual key
	AdminID    int       `json:"admin_id" gorm:"not null"`      // Admin who created the key
	Redeemed   bool      `json:"redeemed" gorm:"default:false"` // If the key has been redeemed
	UserID     int       `json:"user_id"`                       // User who redeemed the key
	
	RedeemedAt time.Time // When the key was redeemed
	CreatedAt  time.Time // When the key was created
	UpdatedAt  time.Time // When the key was last updated
}
