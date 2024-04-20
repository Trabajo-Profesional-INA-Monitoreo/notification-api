package endpoints

import (
	"github.com/Trabajo-Profesional-INA-Monitoreo/notification-api/controllers"
	"github.com/Trabajo-Profesional-INA-Monitoreo/notification-api/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func SetNotificationEndpoints(apiGroup *gin.RouterGroup) {
	controller := controllers.NewNotificationController(services.NewNotificationService())
	testApi := apiGroup.Group("/notificacion")
	{
		testApi.POST("", controller.CreateNotification)
	}
	log.Infof("Configured notification endpoints")
}
