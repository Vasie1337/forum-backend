package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`            // Primary key
	Username string `json:"username" gorm:"unique;not null"` // Username
	Password string `json:"password" gorm:"not null"`        // Password
	Email    string `json:"email" gorm:"unique;not null"`    // Email
}

// MYSQL:
// CREATE TABLE users (
// 	id INT AUTO_INCREMENT PRIMARY KEY,
// 	username VARCHAR(255) NOT NULL UNIQUE,
// 	password VARCHAR(255) NOT NULL,
// 	email VARCHAR(255) NOT NULL UNIQUE,
// );

// TESTS:
// INSERT INTO users (username, password, email) VALUES ('test', 'test', 'test');
