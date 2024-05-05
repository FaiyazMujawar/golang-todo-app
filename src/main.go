package main

import (
	"log"

	"github.com/FaiyazMujawar/golang-todo-app/src/initializers"
	"github.com/FaiyazMujawar/golang-todo-app/src/loaders"
	"github.com/joho/godotenv"
)

func init() {
	log.Default().Println("Initializing App...")
	if godotenv.Load() != nil {
		panic("Could not load .env file")
	}
	initializers.ConnectDatabase()
	initializers.UpdateSchema()
}

func main() {
	app := loaders.GetApp()
	log.Default().Println("Starting Server...")
	app.Run()
}
