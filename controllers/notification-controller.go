package controllers

import (
	"github.com/Trabajo-Profesional-INA-Monitoreo/notification-api/dto"
	"github.com/Trabajo-Profesional-INA-Monitoreo/notification-api/services"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"net/http"
)

type NotificationController interface {
	CreateNotification(ctx *gin.Context)
}

type notificationController struct {
	notificationService services.NotificationService
}

// CreateNotification godoc
//
//			@Summary		Endpoint para crear una notificacion
//			@Tags           Notificacion
//			@Produce		json
//	   		@Param          notification  body  dto.Notification    true    "Notification"
//			@Success		201
//		    @Failure        400 {object} dto.ErrorResponse
//			@Router			/notificacion [post]
func (n notificationController) CreateNotification(ctx *gin.Context) {

	var notification dto.Notification

	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	err := n.notificationService.CreateNotification(notification)

	if err != nil {
		log.Errorf("Error sending notif: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, "Notification created")
}

func NewNotificationController(notificationService services.NotificationService) NotificationController {
	return &notificationController{notificationService}
}
