package services

import (
	"github.com/Trabajo-Profesional-INA-Monitoreo/notification-api/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	NotificationService interface {
		CreateNotification(notification dto.Notification) error
	}
)

type notificationService struct {
}

func (n notificationService) CreateNotification(notification dto.Notification) error {
	bot, err := tgbotapi.NewBotAPI("token")
	if err != nil {
		return err
	}
	chatId := -4114307765
	msg := tgbotapi.NewMessage(int64(chatId), notification.Message)
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func NewNotificationService() NotificationService {
	return &notificationService{}
}
