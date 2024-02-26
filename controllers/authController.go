package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lordofthemind/hostelGinPgGormDocker/helpers"
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
	if !helpers.IsUnique("email", superAdmin.Email) {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	// Check if the phone number is unique
	if !helpers.IsUnique("phone", superAdmin.Phone) {
		c.JSON(http.StatusConflict, gin.H{"error": "Phone number already exists"})
		return
	}

	// Check if the username is unique
	if !helpers.IsUnique("username", superAdmin.Username) {
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
	superAdmin.PasswordHash = string(hashedPassword)

	result := initializers.DB.Create(&superAdmin)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"super-admin": superAdmin, "message": "New super admin created successfully"})
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

	// Check in SuperAdminModel
	var superAdmin models.SuperAdminModel
	resultSuperAdmin := initializers.DB.
		Where("username = ? OR email = ? OR phone = ?", loginCredentials.LoginIdentifier, loginCredentials.LoginIdentifier, loginCredentials.LoginIdentifier).
		First(&superAdmin)
	if resultSuperAdmin.Error == nil {
		// Compare the password
		err := bcrypt.CompareHashAndPassword([]byte(superAdmin.PasswordHash), []byte(loginCredentials.Password))
		if err == nil {
			// generate the token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"subject": superAdmin.ID,
				"expires": time.Now().Add(time.Hour * 24 * 30).Unix(),
			})

			tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while generating token"})
				return
			}

			// send the token or set the cookie
			c.SetSameSite(http.SameSiteLaxMode)
			c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
			c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "user_type": "super_admin"})
			return
		}
	}

	// Check in WardenModel
	var warden models.WardenModel
	resultWarden := initializers.DB.
		Where("username = ? OR email = ? OR phone = ?", loginCredentials.LoginIdentifier, loginCredentials.LoginIdentifier, loginCredentials.LoginIdentifier).
		First(&warden)
	if resultWarden.Error == nil {
		// Compare the password
		err := bcrypt.CompareHashAndPassword([]byte(warden.PasswordHash), []byte(loginCredentials.Password))
		if err == nil {
			// generate the token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"subject": warden.ID,
				"expires": time.Now().Add(time.Hour * 24 * 30).Unix(),
			})

			tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while generating token"})
				return
			}

			// send the token or set the cookie
			c.SetSameSite(http.SameSiteLaxMode)
			c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
			c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "user_type": "warden"})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{"message": user})
}
