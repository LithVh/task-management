package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup, controller *Controller,
	 authMw gin.HandlerFunc, rateLimitMw gin.HandlerFunc) {
	users := group.Group("/users")
	users.Use(authMw) 
	users.Use(rateLimitMw)
	{
		users.GET("/me", controller.MyProfile)
		users.PUT("/me", controller.UpdateMyProfile)
	}
}
