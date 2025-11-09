package middleware

import (
	"task-management/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(config *config.Config) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()

	if config.CORS.Origin == "*" {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowOrigins = []string{config.CORS.Origin}
	}

	if config.Server.Env == "development" {
		corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTION"}
		corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
		corsConfig.ExposeHeaders = []string{"Content-Length"}
		corsConfig.AllowCredentials = true
	}

	return cors.New(corsConfig)
}
