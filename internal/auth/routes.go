package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup, controller *AuthController) {
	auth := group.Group("/auth")
	auth.POST("/register", controller.Register)
	auth.POST("/login", controller.Login)
}
