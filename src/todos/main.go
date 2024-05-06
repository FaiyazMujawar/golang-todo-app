package todos

import (
	"github.com/FaiyazMujawar/golang-todo-app/src/middlewares"
	"github.com/gin-gonic/gin"
)

func AddTodoRouteHandlers(router *gin.Engine) {
	todoRoutes := router.Group("/api/todos", middlewares.ValidateToken)
	{
		todoRoutes.GET("/", getAllTodos)
		todoRoutes.GET("/:id", getTodoById)
		todoRoutes.POST("/", createTodo)
		todoRoutes.PATCH("/:id/mark-done", markDone)
	}
}
