package config

import "time"

var (
	JWTSecretKeyAdmin = []byte("1337")
	JWTSecretKeyUser  = []byte("7331")

	MaxTimeWindowJWTUser  = 1 * time.Minute // Token is only for accessing user resources, so it has a short expiration time
	MaxTimeWindowJWTAdmin = 3 * time.Hour   // Token is used for user management, so it has a longer expiration time

	DSN = "root:@tcp(127.0.0.1:3306)/test-backend?charset=utf8mb4&parseTime=True&loc=Local"

	GinMode = "release"
)
