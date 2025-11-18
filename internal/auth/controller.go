package auth

import (
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service AuthService
}

func NewAuthController(service AuthService) *AuthController {
	return &AuthController{service: service}
}

func (controller *AuthController) Register(c *gin.Context) {
	var dto RegisterRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}

	res, err := controller.service.Register(c.Request.Context(), &dto)
	if err != nil {
		if err.Error() == "email already registered" {
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

	c.IndentedJSON(201, res)
}

func (controller *AuthController) Login(c *gin.Context) {
	var dto LoginRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}

	res, err := controller.service.Login(c.Request.Context(), &dto)
	if err != nil {
		c.IndentedJSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, res)
}
