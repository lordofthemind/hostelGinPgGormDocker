package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToPostgresql() {
	var err error
	postgresqlConnectionString := os.Getenv("ELEPHANT_CONNECTION")
	DB, err = gorm.Open(postgres.Open(postgresqlConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
}
