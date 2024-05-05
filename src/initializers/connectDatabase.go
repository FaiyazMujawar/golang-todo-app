package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := os.Getenv("dsn")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %s", err.Error()))
	}

	log.Default().Println("Connected to Database")
}
