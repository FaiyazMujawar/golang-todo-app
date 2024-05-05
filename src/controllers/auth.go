package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/FaiyazMujawar/golang-todo-app/src/initializers"
	"github.com/FaiyazMujawar/golang-todo-app/src/models"
	jwtService "github.com/FaiyazMujawar/golang-todo-app/src/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

	result := initializers.DB.Create(&user)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Email %s already used", user.Email),
		})
		return
	}
	ctx.AbortWithStatus(http.StatusCreated)
}

func login(ctx *gin.Context) {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var loginRequest LoginRequest
	ctx.BindJSON(&loginRequest)

	var user models.User
	result := initializers.DB.Where("email=?", loginRequest.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "No user found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect password",
		})
		return
	}
	token, err := jwtService.SignToken(jwt.MapClaims{
		"sub": user.ID,
	})
	if err != nil {
		log.Default().Println(err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Token signing failed",
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func AddAuthRoutes(router *gin.Engine) {
	router.POST("/auth/register", register)
	router.POST("/auth/login", login)
}
