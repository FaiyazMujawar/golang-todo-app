package middlewares

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/FaiyazMujawar/golang-todo-app/src/initializers"
	"github.com/FaiyazMujawar/golang-todo-app/src/models"
	jwtService "github.com/FaiyazMujawar/golang-todo-app/src/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ValidateToken(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	match, err := regexp.MatchString("Bearer .+", token)
	if err != nil || !match {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "Valid token must be provided",
		})
		return
	}
	claims, err := jwtService.VerifyToken(token[7:])
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	userId := claims["sub"].(string)
	var user models.User
	result := initializers.DB.First(&user, "id=?", userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "User not found",
		})
		return
	}
	ctx.Set("user", user)
	ctx.Next()
}
