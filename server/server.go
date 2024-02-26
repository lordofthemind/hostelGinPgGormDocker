package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/hostelGinPgGormDocker/controllers"
	"github.com/lordofthemind/hostelGinPgGormDocker/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToPostgresql()
	initializers.SyncPostgresql()
}

func Run() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.Run("localhost:" + os.Getenv("PORT"))
}
