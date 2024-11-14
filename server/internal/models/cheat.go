package models

import "time"

type Cheat struct {
	ID          int    `json:"id" gorm:"primaryKey"`        // Primary key
	Name        string `json:"name" gorm:"unique;not null"` // Cheat name
	Content     string `json:"content" gorm:"not null"`     // Cheat content
	Description string `json:"description" gorm:"not null"` // Cheat description

	UpdatedAt time.Time // When the cheat was last updated 
	CreatedAt time.Time // When the cheat was created
}

// MYSQL:
// CREATE TABLE cheats (
// 	id INT AUTO_INCREMENT PRIMARY KEY,
// 	name VARCHAR(255) NOT NULL UNIQUE,
// 	content TEXT NOT NULL,
// 	description TEXT NOT NULL,
//  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// );

// TESTS:
// INSERT INTO cheats (name, content, description) VALUES ('test', 'test', 'test');
