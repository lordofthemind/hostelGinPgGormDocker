package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/hostelGinPgGormDocker/helpers"
	"github.com/lordofthemind/hostelGinPgGormDocker/initializers"
	"github.com/lordofthemind/hostelGinPgGormDocker/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateCoordinator(c *gin.Context) {
	// Get SuperAdmin ID from context
	superAdminID, exists := c.Get("superAdminID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "SuperAdmin ID not found in context"})
		return
	}

	// Convert superAdminID to uint
	createdBy, ok := superAdminID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error converting SuperAdmin ID"})
		return
	}

	var newWarden models.CoordinatorModel
	rawJSON, _ := c.GetRawData()
	fmt.Println("Raw JSON Payload:", string(rawJSON))

	// Attempt to bind the JSON payload
	if err := c.ShouldBindJSON(&newWarden); err != nil {
		fmt.Println("Error binding JSON:", err.Error())
		c.JSON(400, gin.H{"error": err, "message": "Invalid JSON payload"})
		return
	}

	// Set the CreatedBy field with the SuperAdmin's ID
	newWarden.CreatedBy = createdBy

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

func GetAllWarden(c *gin.Context) {
	var wardens []models.CoordinatorModel
	result := initializers.DB.Find(&wardens)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"data": wardens})
}

func GetWardenByID(c *gin.Context) {
	var warden models.CoordinatorModel
	result := initializers.DB.First(&warden, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"data": warden})
}

func CreateHostel(c *gin.Context) {
	var newHostel models.HostelModel
	if err := c.ShouldBindJSON(&newHostel); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Validate the model
	// if err := newHostel.Validate(); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// Create the hostel
	result := initializers.DB.Create(&newHostel)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, gin.H{"data": newHostel, "message": "New Hostel created successfully"})
}

func GetAllHostel(c *gin.Context) {
	var hostels []models.HostelModel
	result := initializers.DB.Find(&hostels)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"data": hostels})
}

func GetHostelByID(c *gin.Context) {
	var hostel models.HostelModel
	result := initializers.DB.First(&hostel, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"data": hostel})
}

func UpdateHostelByID(c *gin.Context) {
	var hostel models.HostelModel
	result := initializers.DB.First(&hostel, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	if err := c.ShouldBindJSON(&hostel); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	result = initializers.DB.Save(&hostel)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"data": hostel})
}

func DeleteHostelByID(c *gin.Context) {
	var hostel models.HostelModel
	result := initializers.DB.First(&hostel, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	result = initializers.DB.Delete(&hostel)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Hostel deleted successfully"})
}
