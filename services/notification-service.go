package services

import (
	"github.com/Trabajo-Profesional-INA-Monitoreo/notification-api/config"
	"github.com/Trabajo-Profesional-INA-Monitoreo/notification-api/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	NotificationService interface {
		CreateNotification(notification dto.Notification) error
	}
)

type notificationService struct {
	tokenTelegram string
	chatId        int64
}

func (n notificationService) CreateNotification(notification dto.Notification) error {
	bot, err := tgbotapi.NewBotAPI(n.tokenTelegram)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(n.chatId, notification.Message)
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func NewNotificationService(apiConfig *config.ApiConfig) NotificationService {
	return &notificationService{tokenTelegram: apiConfig.TokenTelegram, chatId: apiConfig.ChatId}
}
