package main

import (
	"log"

	"task-management/internal/auth"
	"task-management/internal/config"
	"task-management/internal/database"
	"task-management/internal/middleware"
	"task-management/internal/project"
	"task-management/internal/task"
	"task-management/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {

	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = database.PostgresConnect(config)
	if err != nil {
		log.Fatal(err)
	}
	db := database.GetDB()

	userRepo := user.NewRepository(db)
	authService := auth.NewAuthService(config, userRepo)
	authController := auth.NewAuthController(authService)

	userService := user.NewService(userRepo)
	userController := user.NewController(userService)

	projectRepo := project.NewRepository(db)
	projectService := project.NewService(projectRepo)
	projectController := project.NewController(projectService)

	taskRepo := task.NewRepository(db)
	taskService := task.NewService(taskRepo, projectRepo)
	taskController := task.NewController(taskService)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware(config))
	group := router.Group("/")
	auth.RegisterRoutes(group, authController)
	user.RegisterRoutes(group, userController, middleware.AuthMiddleware(config))
	project.RegisterRoutes(group, projectController, middleware.AuthMiddleware(config))
	task.RegisterRoutes(group, taskController, middleware.AuthMiddleware(config))

	router.Run(":8080")
}
