package models

import "time"

type Key struct {
	ID         int       `json:"id" gorm:"primaryKey"`          // Primary key
	CheatID    int       `json:"cheat_id" gorm:"not null"`      // Cheat ID
	Key        string    `json:"key" gorm:"unique;not null"`    // Actual key
	AdminID    int       `json:"admin_id" gorm:"not null"`      // Admin who created the key
	Redeemed   bool      `json:"redeemed" gorm:"default:false"` // If the key has been redeemed
	RedeemedAt time.Time `json:"redeemed_at"`                   // Timestamp when the key was redeemed
	ValidUntil time.Time `json:"valid_until"`                   // Timestamp when the key expires
	UserID     int       `json:"user_id"`                       // User who redeemed the key
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// MYSQL:
//CREATE TABLE `keys` (
//    id INT AUTO_INCREMENT PRIMARY KEY,
//    cheat_id INT NOT NULL,
//    `key` VARCHAR(255) NOT NULL UNIQUE,
//    admin_id INT NOT NULL,
//    redeemed BOOLEAN DEFAULT FALSE,
//    redeemed_at TIMESTAMP NULL,
//    valid_until TIMESTAMP NULL,
//    user_id INT,
//    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
//);

// TESTS:
// INSERT INTO `keys` (cheat_id, `key`, admin_id, valid_until) VALUES (1, 'test', 1, DATE_ADD(CURRENT_TIMESTAMP, INTERVAL 1 YEAR));
