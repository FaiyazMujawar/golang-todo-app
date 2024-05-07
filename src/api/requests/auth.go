package requests

import "github.com/FaiyazMujawar/golang-todo-app/src/models"

type RegisterUserRequest struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Email     string `json:"email" gorm:"unique" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (request RegisterUserRequest) ToUser() models.User {
	return models.User{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Email:     request.Email,
		Password:  request.Password,
	}
}
