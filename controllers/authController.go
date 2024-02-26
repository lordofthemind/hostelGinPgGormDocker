package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/hostelGinPgGormDocker/initializers"
	"github.com/lordofthemind/hostelGinPgGormDocker/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Sign up a user
	var superAdmin models.SuperAdminModel
	if err := c.ShouldBindJSON(&superAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Validate the model
	if err := superAdmin.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email is unique
	if !isUnique("email", superAdmin.Email) {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	// Check if the phone number is unique
	if !isUnique("phone", superAdmin.Phone) {
		c.JSON(http.StatusConflict, gin.H{"error": "Phone number already exists"})
		return
	}

	// Check if the username is unique
	if !isUnique("username", superAdmin.Username) {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(superAdmin.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}

	// Create the user
	user := models.SuperAdminModel{
		Username:     superAdmin.Username,
		Email:        superAdmin.Email,
		Phone:        superAdmin.Phone,
		Name:         superAdmin.Name,
		PasswordHash: string(hashedPassword),
		Address:      superAdmin.Address,
		IsActive:     superAdmin.IsActive,
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// isUnique checks if a value is unique in the given column of the SuperAdminModel
func isUnique(column, value string) bool {
	var count int64
	initializers.DB.Model(&models.SuperAdminModel{}).Where(fmt.Sprintf("%s = ?", column), value).Count(&count)
	return count == 0
}
