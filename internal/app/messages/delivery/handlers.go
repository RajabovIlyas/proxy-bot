package delivery

import (
	"github.com/RajabovIlyas/proxy-bot/config"
	"github.com/RajabovIlyas/proxy-bot/internal/app/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
)

type messageHandlers struct {
	cfg       *config.Config
	logger    zerolog.Logger
	bot       *tgbotapi.BotAPI
	messageUC messages.UseCase
}

func (m messageHandlers) Messages(update tgbotapi.Update) {
	m.messageUC.SendMessage(update)
}

func NewMessageHandlers(messageUC messages.UseCase, bot *tgbotapi.BotAPI, cfg *config.Config, logger zerolog.Logger) messages.Handlers {
	return &messageHandlers{messageUC: messageUC, bot: bot, cfg: cfg, logger: logger}
}
