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
	r.POST("/create-warden", middlewares.RequireAuthentication, controllers.CreateWarden)
	r.GET("/get-All-warden", middlewares.RequireAuthentication, controllers.GetAllWarden)
	r.GET("/get-warden/:id", middlewares.RequireAuthentication, controllers.GetWardenByID)
	r.POST("/create-Hostel", middlewares.RequireAuthentication, controllers.CreateHostel)
	r.GET("/get-All-Hostel", middlewares.RequireAuthentication, controllers.GetAllHostel)
	r.GET("/get-Hostel/:id", middlewares.RequireAuthentication, controllers.GetHostelByID)
	r.PUT("/update-Hostel/:id", middlewares.RequireAuthentication, controllers.UpdateHostelByID)
	r.DELETE("/delete-Hostel/:id", middlewares.RequireAuthentication, controllers.DeleteHostelByID)
	r.Run("localhost:" + os.Getenv("PORT"))
}
