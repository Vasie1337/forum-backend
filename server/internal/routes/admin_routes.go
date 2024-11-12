package routes

import (
	"server/internal/handlers"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine, adminService *services.AdminService, authService *services.AuthService) {
	adminHandler := handlers.NewAdminHandler(adminService, authService)

	admin := r.Group("/admin")
	{
		admin.POST("/login", adminHandler.AdminLogin)
		admin.GET("/users", adminHandler.GetAllUsers)
	}
}
