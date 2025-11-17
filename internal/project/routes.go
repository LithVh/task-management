package project

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup, controller *Controller,
	 authMw gin.HandlerFunc, rateLimitMw gin.HandlerFunc) {
	projects := group.Group("/projects")
	projects.Use(authMw) 
	projects.Use(rateLimitMw)
	{
		projects.GET("", controller.List)
		projects.POST("", controller.Create)
		projects.GET("/:id", controller.GetByID)
		projects.PUT("/:id", controller.Update)
		projects.DELETE("/:id", controller.Delete)
		projects.POST("/:id/members", controller.AddUser)
	}
}
