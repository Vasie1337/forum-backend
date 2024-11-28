package routes

import (
	"server/internal/handlers"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func LoaderRoutes(r *gin.Engine, loaderService *services.LoaderService) {
	loaderHandler := handlers.NewLoaderHandler(loaderService)

	loader := r.Group("/loader")
	{
		loader.GET("/socket", loaderHandler.SocketHandler)
	}
}
