package user

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers user-related routes
func RegisterRoutes(group *gin.RouterGroup, controller *Controller, authMw gin.HandlerFunc) {
	users := group.Group("/users")
	users.Use(authMw) // Protected routes
	{
		users.GET("/me", controller.MyProfile)
		users.PUT("/me", controller.UpdateMyProfile)
	}
}
