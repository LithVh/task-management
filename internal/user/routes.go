package user

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers user-related routes
func RegisterRoutes(rg *gin.RouterGroup, ctrl *Controller, authMw gin.HandlerFunc) {
	users := rg.Group("/users")
	users.Use(authMw) // Protected routes
	{
		users.GET("/me", ctrl.GetMe)
		users.PUT("/me", ctrl.UpdateMe)
	}
}
