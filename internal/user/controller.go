package user

import (
	"task-management/internal/middleware"
	"task-management/internal/utils"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

// GetMe handles GET /api/v1/users/me
func (ctrl *Controller) GetMe(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		utils.ErrorJSON(c, 401, "Unauthorized")
		return
	}

	user, err := ctrl.service.GetProfile(userID)
	if err != nil {
		utils.ErrorJSON(c, 404, err.Error())
		return
	}

	utils.SuccessJSON(c, 200, user)
}

// UpdateMe handles PUT /api/v1/users/me
func (ctrl *Controller) UpdateMe(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		utils.ErrorJSON(c, 401, "Unauthorized")
		return
	}

	var dto UpdateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		utils.ErrorJSON(c, 400, "Invalid request body: "+err.Error())
		return
	}

	user, err := ctrl.service.UpdateProfile(userID, &dto)
	if err != nil {
		if err.Error() == "email already in use" {
			utils.ErrorJSON(c, 409, err.Error())
			return
		}
		utils.ErrorJSON(c, 500, err.Error())
		return
	}

	utils.SuccessJSON(c, 200, user)
}
