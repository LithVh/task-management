package subtask

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup, ctrl *Controller, authMiddleware gin.HandlerFunc) {
	// Routes for subtasks under tasks
	taskRoutes := group.Group("/tasks/:id/subtasks")
	taskRoutes.Use(authMiddleware)

	taskRoutes.GET("", ctrl.List)
	taskRoutes.POST("", ctrl.Create)

	// taskRoutes.GET("/:id", ctrl.GetByID)
	taskRoutes.PUT("/:id", ctrl.Update)
	taskRoutes.DELETE("/:id", ctrl.Delete)
	taskRoutes.PATCH("/:id/complete", ctrl.ToggleComplete)

}
