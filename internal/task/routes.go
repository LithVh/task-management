package task

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup, ctrl *Controller, authMw gin.HandlerFunc) {
	projects := group.Group("/projects")
	projects.Use(authMw)
	{
		projects.GET("/:id/tasks", ctrl.List)
		projects.POST("/:id/tasks", ctrl.Create)
	}

	tasks := group.Group("/tasks")
	tasks.Use(authMw)
	{
		tasks.GET("/:id", ctrl.GetByID)
		tasks.PUT("/:id", ctrl.Update)
		tasks.DELETE("/:id", ctrl.Delete)
		tasks.PATCH("/:id/complete", ctrl.ToggleComplete)
	}
}
