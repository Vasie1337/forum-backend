package handlers

import (
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User login",
	})
}

func UserRegister(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User register",
	})
}

func UserProfile(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User profile",
	})
}

func UserUpdateProfile(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User update profile",
	})
}

func UserRedeem(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User redeem",
	})
}
