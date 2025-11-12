package subtask

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

//list all subtasks of a task
func (ctrl *Controller) List(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.IndentedJSON(401, gin.H{"error": "user not authenticated"})
		return
	}

	userUUID := userID.(uuid.UUID)

	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseInt(taskIDStr, 10, 64)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "invalid task ID"})
		return
	}

	status := c.Query("status")
	priority := c.Query("priority")

	var statusPtr, priorityPtr *string
	if status != "" {
		statusPtr = &status
	}
	if priority != "" {
		priorityPtr = &priority
	}

	subtasks, err := ctrl.service.List(userUUID, taskID, statusPtr, priorityPtr)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.IndentedJSON(404, gin.H{"error": err.Error()})
			return
		}
		if strings.Contains(err.Error(), "unauthorized") {
			c.IndentedJSON(403, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, ToSubtaskResponseList(subtasks))
}

//creates a new subtask for a given task
func (ctrl *Controller) Create(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.IndentedJSON(401, gin.H{"error": "user not authenticated"})
		return
	}

	userUUID := userID.(uuid.UUID)

	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseInt(taskIDStr, 10, 64)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "invalid task ID"})
		return
	}

	var req CreateSubtaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	subtask, err := ctrl.service.Create(userUUID, taskID, req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.IndentedJSON(404, gin.H{"error": err.Error()})
			return
		}
		if strings.Contains(err.Error(), "unauthorized") {
			c.IndentedJSON(403, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(201, ToSubtaskResponse(subtask))
}

func (ctrl *Controller) GetByID(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userUUID := userID.(uuid.UUID)

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid subtask ID"})
		return
	}

	subtask, err := ctrl.service.GetByID(userUUID, id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if strings.Contains(err.Error(), "unauthorized") {
			c.IndentedJSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, ToSubtaskResponse(subtask))
}

func (ctrl *Controller) Update(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.IndentedJSON(401, gin.H{"error": "user not authenticated"})
		return
	}

	userUUID := userID.(uuid.UUID)

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "invalid subtask ID"})
		return
	}

	var req UpdateSubtaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subtask, err := ctrl.service.Update(userUUID, id, req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if strings.Contains(err.Error(), "unauthorized") {
			c.IndentedJSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, ToSubtaskResponse(subtask))
}

func (ctrl *Controller) ToggleComplete(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.IndentedJSON(401, gin.H{"error": "user not authenticated"})
		return
	}

	userUUID := userID.(uuid.UUID)

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "invalid subtask ID"})
		return
	}

	subtask, err := ctrl.service.ToggleComplete(userUUID, id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if strings.Contains(err.Error(), "unauthorized") {
			c.IndentedJSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, ToSubtaskResponse(subtask))
}

// Delete deletes a specific subtask
// DELETE /subtasks/:id
func (ctrl *Controller) Delete(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userUUID := userID.(uuid.UUID)

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "invalid subtask ID"})
		return
	}

	err = ctrl.service.Delete(userUUID, id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if strings.Contains(err.Error(), "unauthorized") {
			c.IndentedJSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "subtask deleted successfully"})
}

// GetByAssignedUser retrieves all subtasks assigned to a specific user
// GET /users/:id/subtasks
func (ctrl *Controller) GetByAssignedUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userUUID := userID.(uuid.UUID)

	assignedToStr := c.Param("id")
	assignedTo, err := uuid.Parse(assignedToStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	subtasks, err := ctrl.service.GetByAssignedUser(userUUID, assignedTo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, ToSubtaskResponseList(subtasks))
}
