package messages

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type UseCase interface {
	SetMessage(update tgbotapi.Update)
}
