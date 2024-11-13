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

	token, err := h.AuthService.GenerateUserToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	println("Token: ", token)

	c.JSON(200, gin.H{
		"user":  user,
		"token": token,
	})
}

func (h *UserHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		claims, err := h.AuthService.ValidateUserToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func (h *UserHandler) RedeemKey(c *gin.Context) {
	var redemptionRequest struct {
		KeyValue string `json:"key_value"`
		UserID   int    `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&redemptionRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.UserService.RedeemKey(redemptionRequest.UserID, redemptionRequest.KeyValue)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "key redeemed successfully"})
}
