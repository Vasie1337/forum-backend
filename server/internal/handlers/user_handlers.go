package handlers

import (
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
	AuthService *services.AuthService
}

func NewUserHandler(userService *services.UserService, authService *services.AuthService) *UserHandler {
	return &UserHandler{
		UserService: userService,
		AuthService: authService,
	}
}

func (h *UserHandler) UserLogin(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.AuthService.UserLogin(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid username or password"})
		return
	}

	c.JSON(200, user)
}

func (h *UserHandler) RedeemKey(c *gin.Context) {
	var redemptionRequest struct {
		KeyValue string `json:"key_value"`
	}

	if err := c.ShouldBindJSON(&redemptionRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt("user_id") // Assuming you have middleware to extract the user ID

	err := h.UserService.RedeemKey(userID, redemptionRequest.KeyValue)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "key redeemed successfully"})
}
