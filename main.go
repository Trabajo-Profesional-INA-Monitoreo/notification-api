package main

import (
	"fmt"
	"github.com/Trabajo-Profesional-INA-Monitoreo/notification-api/config"
	"github.com/Trabajo-Profesional-INA-Monitoreo/notification-api/endpoints"
	"github.com/Trabajo-Profesional-INA-Monitoreo/notification-api/middlewares"
	"github.com/gin-gonic/gin"
	"log"
)

// @title			Notification API
// @version		1.0
// @description	This API manages the notifications of the forecast model
func main() {
	apiConfig := config.GetConfig()
	config.InitLogger(apiConfig.LogLevel)
	server := gin.New()

	middlewares.SetUpMiddlewares(server)
	endpoints.SetUpEndpoints(server, apiConfig)

	err := server.Run(fmt.Sprintf(":%v", apiConfig.ServerPort))
	if err != nil {
		log.Fatalf("%v", err)
	}
}
