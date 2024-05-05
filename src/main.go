package main

import (
	"log"

	"github.com/FaiyazMujawar/golang-todo-app/src/router"
)

func main() {
	r := router.Router()
	log.Fatalln(r.Run(":3000"))
}
