package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func SignIn(c *gin.Context) {
	// Sign in a user
	var loginCredentials struct {
		LoginIdentifier string `json:"login_identifier" binding:"required"`
		Password        string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginCredentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Get the user from the database using username, email, or phone number
	var user models.SuperAdminModel
	result := initializers.DB.
		Where("username = ? OR email = ? OR phone = ?", loginCredentials.LoginIdentifier, loginCredentials.LoginIdentifier, loginCredentials.LoginIdentifier).
		First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Compare the password
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginCredentials.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}
	// generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.ID,
		"expires": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while generating token"})
		return
	}

	// send the token
	// c.JSON(http.StatusOK, gin.H{"token": tokenString})
	// send the cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{"message": user})
}
