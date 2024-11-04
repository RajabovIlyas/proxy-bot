package usecase

import (
	"github.com/RajabovIlyas/proxy-bot/config"
	"github.com/RajabovIlyas/proxy-bot/internal/app/messages"
	"github.com/RajabovIlyas/proxy-bot/internal/app/proxy"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
	"time"
)

type messageUC struct {
	bot     *tgbotapi.BotAPI
	cfg     *config.Config
	logger  zerolog.Logger
	proxyUC proxy.UseCase
}

func (m messageUC) SendErrorMessage(update tgbotapi.Update, err error) {
	m.logger.Warn().Err(err)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Incorrect input!")

	if _, err := m.bot.Send(msg); err != nil {
		m.logger.Error().Err(err).Msg("failed to send message")
		return
	}
}

func (m messageUC) SendMessage(update tgbotapi.Update) {
	sendMessages, err := m.proxyUC.GetMessages(update.Message)
	if err != nil {
		m.SendErrorMessage(update, err)
		return
	}
	for _, message := range sendMessages {
		time.Sleep(time.Second)
		if _, err := m.bot.Send(message); err != nil {
			m.logger.Error().Err(err).Msg("failed to send message")
			continue
		}
	}
}

func NewMessageUseCase(bot *tgbotapi.BotAPI, cfg *config.Config, logger zerolog.Logger, proxyUC proxy.UseCase) messages.UseCase {
	return &messageUC{bot: bot, cfg: cfg, logger: logger, proxyUC: proxyUC}
}
