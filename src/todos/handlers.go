package todos

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/FaiyazMujawar/golang-todo-app/src/auth"
	"github.com/FaiyazMujawar/golang-todo-app/src/initializers"
	"github.com/FaiyazMujawar/golang-todo-app/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getAllTodos(ctx *gin.Context) {
	loggedInUser, _ := auth.GetLoggedInUser(ctx)

	var todos []models.Todo
	result := initializers.DB.Where("user_id = ?", loggedInUser.ID).Find(&todos)
	if result.Error != nil {
		log.Default().Println(result.Error)

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(ctx *gin.Context) {
	loggedInUser, _ := auth.GetLoggedInUser(ctx)
	todoId := ctx.Param("id")

	var todo models.Todo
	result := initializers.DB.Where("id = ? AND user_id = ?", todoId, loggedInUser.ID).First(&todo)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "No todo with given ID found",
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": result.Error.Error(),
			})
		}
		return
	}

	ctx.IndentedJSON(http.StatusOK, todo)
}

func createTodo(ctx *gin.Context) {
	loggedInUser, _ := auth.GetLoggedInUser(ctx)

	type CreateTodoRequest struct {
		Title       string     `json:"title"`
		Description *string    `json:"description,omitempty"`
		Expiry      *time.Time `json:"expiry,omitempty"`
	}

	var request CreateTodoRequest
	ctx.BindJSON(&request)

	todo := models.Todo{
		Title:       request.Title,
		Description: request.Description,
		Expiry:      request.Expiry,
		User:        *loggedInUser,
	}
	result := initializers.DB.Create(&todo)

	if result.Error != nil {
		log.Default().Println(result.Error)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, todo)
}

func markDone(ctx *gin.Context) {
	loggedInUser, _ := auth.GetLoggedInUser(ctx)
	todoId := ctx.Param("id")

	var todo models.Todo
	result := initializers.DB.Where("id = ? AND user_id = ?", todoId, loggedInUser.ID).First(&todo)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "No todo with given ID found",
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": result.Error.Error(),
			})
		}
		return
	}
	if todo.IsCompleted {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Already marked done",
		})
		return
	}
	todo.IsCompleted = true
	initializers.DB.Save(&todo)
	ctx.AbortWithStatus(http.StatusOK)
}
