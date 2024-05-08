package notes

import (
	"errors"
	"log"
	"net/http"

	"github.com/FaiyazMujawar/golang-todo-app/src/api/requests"
	"github.com/FaiyazMujawar/golang-todo-app/src/auth"
	"github.com/FaiyazMujawar/golang-todo-app/src/initializers"
	"github.com/FaiyazMujawar/golang-todo-app/src/models"
	"github.com/FaiyazMujawar/golang-todo-app/src/utils"
	"github.com/FaiyazMujawar/golang-todo-app/src/utils/s3"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	loggedInUser, _ := auth.GetLoggedInUser(ctx)

	var request requests.CreateNoteRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		errorMessages := utils.ToErrorMessages(err.(validator.ValidationErrors))
		ctx.AbortWithStatusJSON(400, errorMessages)
		return
	}

	urls := s3.UploadFiles(request.Media...)
	note := request.ToNote(urls, *loggedInUser)

	result := initializers.DB.Create(&note)
	if result.Error != nil {
		log.Default().Println("error creating note: ", result.Error)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, note)
}

func deleteById(ctx *gin.Context) {
	loggedInUser, _ := auth.GetLoggedInUser(ctx)
	noteId := ctx.Param("id")

	var note models.Note
	result := initializers.DB.Where("id = ? AND user_id = ?", noteId, loggedInUser.ID).First(&note)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "No note found",
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": result.Error.Error(),
			})
		}
		return
	}

	// Remove media from S3
	urls := utils.Map(note.Media, func(value string) string { return value })
	urls = utils.Map(urls, extractObjectKeyFromUrl)
	err := s3.DeleteFiles(urls...)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	initializers.DB.Delete(note)
	ctx.AbortWithStatus(http.StatusOK)
}
