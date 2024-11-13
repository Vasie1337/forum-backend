package repository

import (
	"server/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Database structure for the server:
//
// CREATE DATABASE test-backend;
//
// USE test-backend;
//
// CREATE TABLE admins (
// 	id INT AUTO_INCREMENT PRIMARY KEY,
// 	username VARCHAR(255) NOT NULL UNIQUE,
// 	password VARCHAR(255) NOT NULL
// );
//
// CREATE TABLE users (
// 	id INT AUTO_INCREMENT PRIMARY KEY,
// 	username VARCHAR(255) NOT NULL UNIQUE,
// 	password VARCHAR(255) NOT NULL
// );
//
// CREATE TABLE keys (
// 	id INT AUTO_INCREMENT PRIMARY KEY,
// 	value VARCHAR(255) NOT NULL UNIQUE
// );
