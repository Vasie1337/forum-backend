package main

import (
	"server/internal/config"
	"server/internal/repository"
	"server/internal/routes"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.GinMode)

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

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running",
		})
	})

	r.Run(":8080")
}
