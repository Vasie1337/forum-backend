package main

import (
	"server/internal/repository"
	"server/internal/routes"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, err := repository.NewDB()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	keyRepo := repository.NewKeyRepository(db)
	adminRepo := repository.NewAdminRepository(db)

	userService := services.NewUserService(userRepo, keyRepo)
	adminService := services.NewAdminService(userRepo, keyRepo)
	authService := services.NewAuthService(adminRepo, userRepo)

	routes.AdminRoutes(r, adminService, authService)
	routes.UserRoutes(r, userService, authService)

	r.Run(":8080")
}
