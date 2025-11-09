package task

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) List(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.IndentedJSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	userUUID := userID.(uuid.UUID)

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "invalid project ID",
		})
		return
	}

	filters := make(map[string]interface{})
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if priority := c.Query("priority"); priority != "" {
		filters["priority"] = priority
	}

	tasks, err := ctrl.service.List(projectID, userUUID, filters)
	if err != nil {
		if err.Error() == "project not found" {
			c.IndentedJSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err.Error() == "unauthorized: you don't own this project" {
			c.IndentedJSON(403, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, tasks)
}

func (ctrl *Controller) Create(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.IndentedJSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	userUUID := userID.(uuid.UUID)

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "invalid project ID",
		})
		return
	}

	var dto CreateTaskRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	task, err := ctrl.service.Create(projectID, userUUID, &dto)
	if err != nil {
		if err.Error() == "project not found" {
			c.IndentedJSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err.Error() == "unauthorized: you don't own this project" {
			c.IndentedJSON(403, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(201, task)
}

func (ctrl *Controller) GetByID(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.IndentedJSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	userUUID := userID.(uuid.UUID)

	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "invalid task ID",
		})
		return
	}

	task, err := ctrl.service.GetByID(taskID, userUUID)
	if err != nil {
		if err.Error() == "task not found" {
			c.IndentedJSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err.Error() == "unauthorized: you don't own this project" {
			c.IndentedJSON(403, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, task)
}

func (ctrl *Controller) Update(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.IndentedJSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	userUUID := userID.(uuid.UUID)

	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "invalid task ID",
		})
		return
	}

	var dto UpdateTaskRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	task, err := ctrl.service.Update(taskID, userUUID, &dto)
	if err != nil {
		if err.Error() == "task not found" {
			c.IndentedJSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err.Error() == "unauthorized: you don't own this project" {
			c.IndentedJSON(403, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, task)
}

func (ctrl *Controller) ToggleComplete(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.IndentedJSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	userUUID := userID.(uuid.UUID)

	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "invalid task ID",
		})
		return
	}

	task, err := ctrl.service.ToggleComplete(taskID, userUUID)
	if err != nil {
		if err.Error() == "task not found" {
			c.IndentedJSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err.Error() == "unauthorized: you don't own this project" {
			c.IndentedJSON(403, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, task)
}

func (ctrl *Controller) Delete(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.IndentedJSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	userUUID := userID.(uuid.UUID)

	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "invalid task ID",
		})
		return
	}

	err = ctrl.service.Delete(taskID, userUUID)
	if err != nil {
		if err.Error() == "task not found" {
			c.IndentedJSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err.Error() == "unauthorized: you don't own this project" {
			c.IndentedJSON(403, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "Task deleted successfully"})
}
