package main

import (
	"github.com/FaiyazMujawar/golang-todo-app/src/loaders"
)

func main() {
	app := loaders.GetApp()
	app.Run()
}
