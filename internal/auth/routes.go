package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup, controller *AuthController, rateLimit gin.HandlerFunc) {
	auth := group.Group("/auth")
	auth.Use(rateLimit)
	auth.POST("/register", controller.Register)
	auth.POST("/login", controller.Login)
}
