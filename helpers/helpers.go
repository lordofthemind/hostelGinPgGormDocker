package helpers

import (
	"fmt"

	"github.com/lordofthemind/hostelGinPgGormDocker/initializers"
	"github.com/lordofthemind/hostelGinPgGormDocker/models"
)

// isUnique checks if a value is unique in the given column of the SuperAdminModel
func IsUnique(column, value string) bool {
	var count int64
	initializers.DB.Model(&models.SuperAdminModel{}).Where(fmt.Sprintf("%s = ?", column), value).Count(&count)
	return count == 0
}
