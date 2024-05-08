package requests

import (
	"mime/multipart"

	"github.com/FaiyazMujawar/golang-todo-app/src/models"
)

type CreateNoteRequest struct {
	Title string                  `form:"title" binding:"required"`
	Body  string                  `form:"body"`
	Media []*multipart.FileHeader `form:"media"`
}

func (request *CreateNoteRequest) ToNote(media []string, user models.User) models.Note {
	return models.Note{
		Title:  request.Title,
		Body:   request.Body,
		Media:  media,
		UserID: user.ID,
		User:   user,
	}
}
