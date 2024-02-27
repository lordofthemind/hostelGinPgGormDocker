package initializers

import (
	"fmt"

	"github.com/lordofthemind/hostelGinPgGormDocker/models"
)

func SyncPostgresql() {
	DB.AutoMigrate(&models.StudentModel{})
	fmt.Println("Synchronized the student model")
	DB.AutoMigrate(&models.BedModel{})
	fmt.Println("Synchronized the bed model")
	DB.AutoMigrate(&models.HostelModel{})
	fmt.Println("Synchronized the hostel model")
	DB.AutoMigrate(&models.CoordinatorModel{})
	fmt.Println("Synchronized the coordinator model")
	DB.AutoMigrate(&models.SuperAdminModel{})
	fmt.Println("Synchronized the super admin model")

	// DB.AutoMigrate(&models.StudentModel{}, &models.BedModel{}, &models.HostelModel{}, &models.CoordinatorModel{}, &models.SuperAdminModel{})

	fmt.Println("Synchronized the database")
}
