package initializers

import "github.com/FaiyazMujawar/golang-todo-app/src/models"

func UpdateSchema() {
	DB.AutoMigrate(&models.User{}, &models.Todo{})
}
