package handlers

import (
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminService *services.AdminService
	AuthService  *services.AuthService
}

func NewAdminHandler(adminService *services.AdminService, authService *services.AuthService) *AdminHandler {
	return &AdminHandler{
		AdminService: adminService,
		AuthService:  authService,
	}
}

func (h *AdminHandler) AdminLogin(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	admin, err := h.AuthService.AdminLogin(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid username or password"})
		return
	}

	c.JSON(200, admin)
}

func (h *AdminHandler) GetAllUsers(c *gin.Context) {
	users, err := h.AdminService.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}
