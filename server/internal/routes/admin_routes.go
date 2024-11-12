package routes

import (
	"server/internal/handlers"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		// Admin routes
		admin.POST("/login", handlers.AdminLogin)

		// User routes
		admin.GET("/users", handlers.GetAllUsers)
		admin.POST("/users", handlers.CreateUser)
		admin.GET("/users/:id", handlers.GetUserByID)
		admin.PUT("/users/:id", handlers.UpdateUser)
		admin.DELETE("/users/:id", handlers.DeleteUser)

		// Key routes
		admin.GET("/keys", handlers.GetAllKeys)
		admin.POST("/keys", handlers.CreateKey)
		admin.GET("/keys/:id", handlers.GetKeyByID)
		admin.PUT("/keys/:id", handlers.UpdateKey)
		admin.DELETE("/keys/:id", handlers.DeleteKey)

		// Cheats routes
		admin.GET("/cheats", handlers.GetAllCheats)
		admin.POST("/cheats", handlers.CreateCheat)
		admin.GET("/cheats/:id", handlers.GetCheatByID)
		admin.PUT("/cheats/:id", handlers.UpdateCheat)
		admin.DELETE("/cheats/:id", handlers.DeleteCheat)
	}

}
