package initializers

import (
	"fmt"

	"github.com/lordofthemind/hostelGinPgGormDocker/models"
)

func SyncPostgresql() {
	// DB.AutoMigrate(&models.SuperAdminModel{}, &models.AdminModel{}, &models.StudentModel{}, &models.HostelModel{}, &models.RoomModel{}, &models.BedModel{}, &models.Visitor)
	DB.AutoMigrate(&models.SuperAdminModel{})
	DB.AutoMigrate(&models.WardenModel{})
	DB.AutoMigrate(&models.StudentModel{})
	DB.AutoMigrate(&models.HostelModel{})

	fmt.Println("Synchronized the database")
}
