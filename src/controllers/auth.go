package auth

import (
	"net/http"

	"github.com/FaiyazMujawar/golang-todo-app/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func register(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	user.Id = uuid.New().String()
	ctx.IndentedJSON(http.StatusOK, user)
}

func AddAuthRoutes(router *gin.Engine) {
	router.POST("/auth/register", register)
}
