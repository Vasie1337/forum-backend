package models

type Admin struct {
	ID       int    `json:"id" gorm:"primaryKey"`            // Primary key
	Username string `json:"username" gorm:"unique;not null"` // Username
	Password string `json:"password" gorm:"not null"`        // Password
}

// MYSQL:
// CREATE TABLE admins (
// 	id INT AUTO_INCREMENT PRIMARY KEY,
// 	username VARCHAR(255) NOT NULL UNIQUE,
// 	password VARCHAR(255) NOT NULL
// );

// TESTS:
// INSERT INTO admins (username, password) VALUES ('test', 'test');
