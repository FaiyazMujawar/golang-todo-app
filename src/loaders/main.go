package loaders

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type App struct {
	port   int16
	router *gin.Engine
}

func (app *App) Run() {
	if app.router == nil {
		panic("Router must be initialized")
	}
	log.Fatalln(app.router.Run(fmt.Sprintf(":%v", app.port)))
}

func GetApp() App {
	router := Router()

	return App{
		port:   3000,
		router: router,
	}
}
