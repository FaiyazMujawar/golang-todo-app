package loaders

import (
	"net/http"

	auth "github.com/FaiyazMujawar/golang-todo-app/src/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/api", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	auth.AddAuthRoutes(router)

	return router
}
