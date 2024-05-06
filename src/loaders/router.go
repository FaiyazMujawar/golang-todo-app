package loaders

import (
	"net/http"

	"github.com/FaiyazMujawar/golang-todo-app/src/auth"
	"github.com/FaiyazMujawar/golang-todo-app/src/middlewares"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/api", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	router.GET("/protected", middlewares.ValidateToken, func(ctx *gin.Context) {
		ctx.String(200, "PROTECTED")
	})

	auth.AddAuthRoutes(router)

	return router
}
