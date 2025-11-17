package subtask

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup, ctrl *Controller,
	 authMiddleware gin.HandlerFunc, rateLimitMw gin.HandlerFunc) {
	taskRoutes := group.Group("tasks/:id/subtasks")
	taskRoutes.Use(authMiddleware)
	taskRoutes.Use(rateLimitMw)

	taskRoutes.GET("", ctrl.List)
	taskRoutes.POST("", ctrl.Create)

	// taskRoutes.GET("/:id", ctrl.GetByID)
	taskRoutes.PUT("/:id", ctrl.Update)
	taskRoutes.DELETE("/:id", ctrl.Delete)
	taskRoutes.PATCH("/:id/complete", ctrl.ToggleComplete)

}
