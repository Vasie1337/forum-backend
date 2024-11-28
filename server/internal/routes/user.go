package routes

import (
	"server/internal/handlers"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userService *services.UserService, authService *services.AuthService) {
	userHandler := handlers.NewUserHandler(userService, authService)

	user := r.Group("/user")
	{
		user.POST("/login", userHandler.UserLogin)
		user.GET("/redeem-key", userHandler.AuthMiddleware(), userHandler.RedeemKey)
	}
}
