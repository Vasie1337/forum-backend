package handlers

import (
	"server/internal/services"
	"strconv"

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

	token, err := h.AuthService.GenerateAdminToken(admin.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(200, gin.H{
		"admin": admin,
		"token": token,
	})
}

func (h *AdminHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		println("AuthMiddleware Token: ", tokenString)
		claims, err := h.AuthService.ValidateAdminToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func (h *AdminHandler) GetAllUsers(c *gin.Context) {
	users, err := h.AdminService.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}

func (h *AdminHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user ID"})
		return
	}
	user, err := h.AdminService.GetUserByID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}
