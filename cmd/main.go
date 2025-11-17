package main

import (
	"log"
	"time"

	"task-management/internal/auth"
	"task-management/internal/config"
	"task-management/internal/database"
	"task-management/internal/middleware"
	"task-management/internal/project"
	"task-management/internal/subtask"
	"task-management/internal/task"
	"task-management/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {

	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	postgres, err := database.PostgresConnect(config)
	if err != nil {
		log.Fatal(err)
	}

	redis, err := database.RedisConnect(config)
	if err != nil {
		log.Fatal(err)
	}

	// db := database.GetDB()

	userRepo := user.NewRepository(postgres)
	authService := auth.NewAuthService(config, userRepo)
	authController := auth.NewAuthController(authService)

	userService := user.NewService(userRepo)
	userController := user.NewController(userService)

	projectRepo := project.NewRepository(postgres)
	projectService := project.NewService(projectRepo)
	projectController := project.NewController(projectService)

	taskRepo := task.NewRepository(postgres)
	taskService := task.NewService(taskRepo, projectRepo)
	taskController := task.NewController(taskService)

	subtaskRepo := subtask.NewRepository(postgres)
	subtaskService := subtask.NewService(subtaskRepo, taskService)
	subtaskController := subtask.NewController(subtaskService)

	postLoggedIn := middleware.RateLimiterMiddleware(*redis, 1000, 3 * time.Minute)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware(config))
	router.Use(middleware.TimeoutMiddleware(10 * time.Second))
	group := router.Group("/")
	auth.RegisterRoutes(group, authController, middleware.RateLimiterMiddleware(*redis, 10, 60 * time.Second))
	
	user.RegisterRoutes(group, userController, middleware.AuthMiddleware(config), postLoggedIn)
	project.RegisterRoutes(group, projectController, middleware.AuthMiddleware(config), postLoggedIn)
	task.RegisterRoutes(group, taskController, middleware.AuthMiddleware(config), postLoggedIn)
	subtask.RegisterRoutes(group, subtaskController, middleware.AuthMiddleware(config), postLoggedIn)

	router.Run(":8080")
}
