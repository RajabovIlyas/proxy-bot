package proxy

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type UseCase interface {
	GetMessages(*tgbotapi.Message) ([]tgbotapi.MessageConfig, error)
}
