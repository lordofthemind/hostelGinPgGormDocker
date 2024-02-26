package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/hostelGinPgGormDocker/controllers"
	"github.com/lordofthemind/hostelGinPgGormDocker/initializers"
	"github.com/lordofthemind/hostelGinPgGormDocker/middlewares"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToPostgresql()
	initializers.SyncPostgresql()
}

func Run() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)
	r.GET("/validate", middlewares.RequireAuthentication, controllers.Validate)
	r.Run("localhost:" + os.Getenv("PORT"))
}
