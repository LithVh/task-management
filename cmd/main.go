package main

import (
	"log"

	"task-management/internal/config"
	"task-management/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {

	config := config.Load()

	err := database.PostgresConnect(config)
	if err != nil {
		log.Fatal(err)
	}

	db := database.GetDB()

	router := gin.Default()

	router.Run(":8080")

}
