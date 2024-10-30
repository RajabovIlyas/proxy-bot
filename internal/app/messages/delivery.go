package messages

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Handlers interface {
	InitMessageTGBot()
	Messages(tgbotapi.Update)
}
