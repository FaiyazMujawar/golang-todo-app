package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/FaiyazMujawar/golang-todo-app/src/api/requests"
	"github.com/FaiyazMujawar/golang-todo-app/src/initializers"
	"github.com/FaiyazMujawar/golang-todo-app/src/models"
	"github.com/FaiyazMujawar/golang-todo-app/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func register(ctx *gin.Context) {
	var request requests.RegisterUserRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		errorMessages := utils.ToErrorMessages(err.(validator.ValidationErrors))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Data",
			"errors":  errorMessages,
		})
		return
	}

	user := request.ToUser()
	result := initializers.DB.Create(&user)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Email %s already used", user.Email),
		})
		return
	}
	ctx.AbortWithStatus(http.StatusCreated)
}

func login(ctx *gin.Context) {

	var loginRequest requests.LoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		errorMessages := utils.ToErrorMessages(err.(validator.ValidationErrors))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Data",
			"errors":  errorMessages,
		})
		return
	}

	var user models.User
	result := initializers.DB.Where("email=?", loginRequest.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "No user found",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect password",
		})
		return
	}
	token, err := SignToken(jwt.MapClaims{
		"sub": user.ID,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Token signing failed",
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"token": token,
	})
}
