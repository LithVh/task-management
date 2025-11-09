package project

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup, controller *Controller, authMw gin.HandlerFunc) {
	projects := group.Group("/projects")
	projects.Use(authMw) 
	{
		projects.GET("", controller.List)
		projects.POST("", controller.Create)
		projects.GET("/:id", controller.GetByID)
		projects.PUT("/:id", controller.Update)
		projects.DELETE("/:id", controller.Delete)
	}
}
