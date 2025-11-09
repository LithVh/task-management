package middleware

import (
	"strings"

	"task-management/internal/config"
	"task-management/internal/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "authorization header required",
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "invalid authorization format",
			})
			return
		}

		token, err := utils.ParseToken(parts[1], *config)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "token invalid or expired",
			})
			return
		}

		c.Set("userID", token.UserID)
		c.Next()

	}
}



// // GetUserID extracts user_id from gin context
// func GetUserID(c *gin.Context) (uuid.UUID, bool) {
// 	userID, exists := c.Get("user_id")
// 	if !exists {
// 		return uuid.Nil, false
// 	}

// 	id, ok := userID.(uuid.UUID)
// 	return id, ok
// }
