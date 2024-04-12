package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lordofthemind/hostelGinPgGormDocker/initializers"
	"github.com/lordofthemind/hostelGinPgGormDocker/models"
)

func RequireAuthentication(c *gin.Context) {
	fmt.Println("Required Authentication middleware")

	// Get the cookie off request
	tokenString, err := c.Cookie("SuperAdminAuthorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Check if the cookie is valid
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiration date
		if float64(time.Now().Unix()) > claims["expires"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Find the user with token subject
		var user models.SuperAdminModel
		result := initializers.DB.First(&user, "id = ?", claims["subject"])
		if result.Error != nil || result.RowsAffected == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach SuperAdmin ID to the request context
		c.Set("superAdminID", user.ID)

		// Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
