package main

import (
	"log"

	"request-mapper/api/controller"
	"request-mapper/api/repository"
	"request-mapper/api/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// hanndle cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	repository := repository.NewRequestMapperRepository()
	service := service.NewRequestMapperService(repository)
	controller := controller.NewRequestMapperController(service)

	// setup routes
	api := router.Group("/api")
	{
		api.POST("/map-request", controller.MapRequest)
	}

	// start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
