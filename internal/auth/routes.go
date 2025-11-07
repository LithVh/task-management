package auth

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers auth-related routes
func RegisterRoutes(rg *gin.RouterGroup, ctrl *Controller) {
	auth := rg.Group("/auth")
	auth.POST("/register", ctrl.Register)
	auth.POST("/login", ctrl.Login)
	auth.POST("/logout", ctrl.Logout)
}
