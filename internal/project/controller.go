package project

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

func (controller *Controller) List(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.IndentedJSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userUUID := userID.(uuid.UUID)

	projects, err := controller.service.List(userUUID)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, projects)
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

	var dto CreateProjectRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	project, err := ctrl.service.Create(userUUID, &dto)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(201, project)
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

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "invalid project ID",
		})
		return
	}

	project, err := ctrl.service.GetByID(projectID, userUUID)
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

	c.IndentedJSON(200, project)
}

func (controller *Controller) Update(c *gin.Context) {
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
		c.IndentedJSON(400, gin.H{"error": "invalid project id"})
		return
	}

	var dto UpdateProjectRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.IndentedJSON(400, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	project, err := controller.service.Update(projectID, userUUID, &dto)
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

	c.IndentedJSON(200, project)
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

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.service.Delete(projectID, userUUID)
	if err != nil {
		if err.Error() == "project not found" {
			c.IndentedJSON(404, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "unauthorized: you don't own this project" {
			c.IndentedJSON(403, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "Project deleted successfully"})
}
