package routes

import (
	"server/internal/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/login", handlers.UserLogin)
		user.POST("/register", handlers.UserRegister)
		user.GET("/profile", handlers.UserProfile)
		user.PUT("/profile", handlers.UserUpdateProfile)
		user.POST("/redeem", handlers.UserRedeem)
	}
}
