package notes

import (
	"errors"
	"net/http"

	"github.com/FaiyazMujawar/golang-todo-app/src/auth"
	"github.com/FaiyazMujawar/golang-todo-app/src/initializers"
	"github.com/FaiyazMujawar/golang-todo-app/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getAllNotes(ctx *gin.Context) {
	loggedInUser, _ := auth.GetLoggedInUser(ctx)

	var notes []models.Note
	result := initializers.DB.Where("user_id = ?", loggedInUser.ID).Find(&notes)
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, notes)
}

func getNoteById(ctx *gin.Context) {
	loggedInUser, _ := auth.GetLoggedInUser(ctx)
	noteId := ctx.Param("id")

	var note models.Note
	result := initializers.DB.Where("id = ? AND user_id = ?", noteId, loggedInUser.ID).First(&note)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Note not found",
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": result.Error.Error(),
			})
		}
		return
	}

	ctx.IndentedJSON(http.StatusOK, note)
}

func create(ctx *gin.Context) {
	// TODO: create note
}
