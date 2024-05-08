package initializers

import (
	"fmt"
	"log"

	"github.com/FaiyazMujawar/golang-todo-app/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(postgres.Open(config.Dsn()), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %s", err.Error()))
	}

	log.Default().Println("Connected to Database")
}
