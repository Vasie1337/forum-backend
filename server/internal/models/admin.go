package models

import "time"

type Admin struct {
	ID        int       `json:"id" gorm:"primaryKey"`            // Primary key
	Username  string    `json:"username" gorm:"unique;not null"` // Username
	Password  string    `json:"password" gorm:"not null"`        // Password
	
	CreatedAt time.Time // When the admin was created
	UpdatedAt time.Time // When the admin was last updated
}
