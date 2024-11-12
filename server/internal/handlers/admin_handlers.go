package handlers

import "github.com/gin-gonic/gin"

func AdminLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Admin login",
	})
}

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all users",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create user",
	})
}

func GetUserByID(c *gin.Context) {
}

func UpdateUser(c *gin.Context) {
}

func DeleteUser(c *gin.Context) {
}

func GetAllKeys(c *gin.Context) {
}

func CreateKey(c *gin.Context) {
}

func GetKeyByID(c *gin.Context) {
}

func UpdateKey(c *gin.Context) {
}

func DeleteKey(c *gin.Context) {
}

func GetAllCheats(c *gin.Context) {
}

func CreateCheat(c *gin.Context) {
}

func GetCheatByID(c *gin.Context) {
}

func UpdateCheat(c *gin.Context) {
}

func DeleteCheat(c *gin.Context) {
}
