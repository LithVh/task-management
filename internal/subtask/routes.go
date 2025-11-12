package subtask

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup, ctrl *Controller, authMiddleware gin.HandlerFunc) {
	// Routes for subtasks under tasks
	taskRoutes := group.Group("/tasks/:id/subtasks")
	taskRoutes.Use(authMiddleware)
	{
		taskRoutes.GET("", ctrl.List)
		taskRoutes.POST("", ctrl.Create)
	}

	// Routes for individual subtasks
	subtaskRoutes := group.Group("/subtasks")
	subtaskRoutes.Use(authMiddleware)
	{
		subtaskRoutes.GET("/:id", ctrl.GetByID)
		subtaskRoutes.PUT("/:id", ctrl.Update)
		subtaskRoutes.DELETE("/:id", ctrl.Delete)
		subtaskRoutes.PATCH("/:id/complete", ctrl.ToggleComplete)
	}

	// Routes for user-assigned subtasks
	userRoutes := group.Group("/users/:id/subtasks")
	userRoutes.Use(authMiddleware)
	{
		userRoutes.GET("", ctrl.GetByAssignedUser)
	}
}
