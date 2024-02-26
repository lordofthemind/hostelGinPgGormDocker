package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/hostelGinPgGormDocker/helpers"
	"github.com/lordofthemind/hostelGinPgGormDocker/initializers"
	"github.com/lordofthemind/hostelGinPgGormDocker/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateWarden(c *gin.Context) {
	var newWarden models.WardenModel
	if err := c.ShouldBindJSON(&newWarden); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Validate the model
	if err := newWarden.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email is unique
	if !helpers.IsUnique("email", newWarden.Email) {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	// Check if the phone number is unique
	if !helpers.IsUnique("phone", newWarden.Phone) {
		c.JSON(http.StatusConflict, gin.H{"error": "Phone number already exists"})
		return
	}

	// Check if the username is unique
	if !helpers.IsUnique("username", newWarden.Username) {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newWarden.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}

	// Create the user
	newWarden.PasswordHash = string(hashedPassword)

	result := initializers.DB.Create(&newWarden)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, gin.H{"data": newWarden, "message": "New Warden account created successfully"})
}
