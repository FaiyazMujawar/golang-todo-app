package auth

import (
	"fmt"

	"github.com/FaiyazMujawar/golang-todo-app/src/models"
	"github.com/gin-gonic/gin"
)

// Returns logged in user set in the context
func GetLoggedInUser(ctx *gin.Context) (*models.User, error) {
	loggedInUser, exists := ctx.Get("user")
	if !exists {
		return nil, fmt.Errorf("no user logged in")
	}
	user := loggedInUser.(models.User)
	return &user, nil
}
