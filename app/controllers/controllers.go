package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/znxxx/Pitchaya_agnos_backend/services"
)

func HandlePasswordSubmission(c *gin.Context) {
	var requestBody map[string]string
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	password, exists := requestBody["init_password"]
	if !exists {
		c.JSON(400, gin.H{"error": "No init_password in request body"})
		return
	} else if len(password) > 40 {
		c.JSON(400, gin.H{"error": "Password length must be between 1 to 40 characters"})
		return
	}

	passwordService := services.NewPasswordService()
	numSteps := passwordService.StrongPasswordSteps(password)

	c.JSON(200, gin.H{"num_of_steps": numSteps})
}
