package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

// GetMe handles GET /api/v1/users/me
func (controller *Controller) MyProfile(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.IndentedJSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userUUID := userID.(uuid.UUID)

	// userUUID, err := uuid.Parse(userID)
	// if err != nil {
	// 	c.IndentedJSON(500, gin.H{
	// 		"error": "id parsing error",
	// 	})
	// 	return
	// }

	user, err := controller.service.GetProfile(userUUID)
	if err != nil {
		c.IndentedJSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, user)
}

func (controller *Controller) UpdateMyProfile(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.IndentedJSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userUUID := userID.(uuid.UUID)

	var dto UpdateUserRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}

	user, err := controller.service.UpdateProfile(userUUID, &dto)
	if err != nil {
		if err.Error() == "email already in use" {
			c.IndentedJSON(409, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, user)
}
