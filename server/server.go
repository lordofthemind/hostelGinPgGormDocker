package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/hostelGinPgGormDocker/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToPostgresql()
	initializers.SyncPostgresql()
}

func Run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:" + os.Getenv("PORT"))
}
