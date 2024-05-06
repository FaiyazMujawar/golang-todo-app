package auth

import "github.com/gin-gonic/gin"

func AddAuthRoutes(router *gin.Engine) {
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", register)
		authRoutes.POST("/login", login)
	}
}
