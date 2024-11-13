package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/test-backend?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
