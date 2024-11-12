package main

import (
	"server/internal/repository"
	"server/internal/routes"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userRepo := repository.NewUserRepository()
	keyRepo := repository.NewKeyRepository()
	
	userService := services.NewUserService(userRepo, keyRepo)
	adminService := services.NewAdminService()
	authService := services.NewAuthService()

	routes.AdminRoutes(r, adminService, authService)
	routes.UserRoutes(r, userService, authService)

	r.Run(":8080")
}
