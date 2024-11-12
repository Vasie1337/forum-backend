package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`            // Primary key
	Username  string    `json:"username" gorm:"unique;not null"` // Username
	Password  string    `json:"password" gorm:"not null"`        // Password
	Email     string    `json:"email" gorm:"unique;not null"`    // Email
	
	CreatedAt time.Time // When the user was created
	UpdatedAt time.Time // When the user was last updated
}
