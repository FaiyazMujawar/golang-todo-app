package notes

import (
	"github.com/FaiyazMujawar/golang-todo-app/src/middlewares"
	"github.com/gin-gonic/gin"
)

func AddNotesRouteHandlers(router *gin.Engine) {
	notesRoutes := router.Group("/api/notes", middlewares.ValidateToken)
	{
		notesRoutes.GET("/", getAllNotes)
		notesRoutes.GET("/:id", getNoteById)
		notesRoutes.POST("/", create)
	}
}
