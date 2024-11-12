package models

import "time"

type Admin struct {
	ID       int    `json:"id" gorm:"primaryKey"`            // Primary key
	Username string `json:"username" gorm:"unique;not null"` // Username
	Password string `json:"password" gorm:"not null"`        // Password

	CreatedAt time.Time // When the admin was created
	UpdatedAt time.Time // When the admin was last updated
}

// MYSQL:
// CREATE TABLE admins (
// 	id INT AUTO_INCREMENT PRIMARY KEY,
// 	username VARCHAR(255) NOT NULL UNIQUE,
// 	password VARCHAR(255) NOT NULL,
//  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// );

// TESTS:
// INSERT INTO admins (username, password) VALUES ('test', 'test');
