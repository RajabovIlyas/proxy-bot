package usecase

import (
	"encoding/json"
	"errors"
	"github.com/RajabovIlyas/proxy-bot/config"
	"github.com/RajabovIlyas/proxy-bot/internal/app/messages"
	"github.com/RajabovIlyas/proxy-bot/internal/app/utils"
	"github.com/RajabovIlyas/proxy-bot/internal/pkg/buttons"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
	"time"
)

type messageUC struct {
	bot    *tgbotapi.BotAPI
	cfg    *config.Config
	logger zerolog.Logger
}

func (m messageUC) SendErrorMessage(update tgbotapi.Update, err error) {
	m.logger.Warn().Err(err)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Incorrect input!")

	if _, err := m.bot.Send(msg); err != nil {
		panic(err)
	}
}

func (m messageUC) SetMessage(update tgbotapi.Update) {
	jsonData, err := json.MarshalIndent(update.Message, "", "  ")
	if err != nil {
		return
	}

	m.logger.Info().Msgf("Message: %v", string(jsonData))

	if update.Message == nil {
		m.SendErrorMessage(update, errors.New("incorrect input"))
		return
	}

	urls := utils.FilterURLs(update.Message)

	if len(urls) == 0 {
		m.SendErrorMessage(update, errors.New("incorrect input"))
		return
	}

	for _, url := range urls {
		m.logger.Info().Msgf("len(%d): %v", len(urls), url)
		proxyParams, err := utils.GetURLParams(url)

		if err != nil {
			continue
		}

		text := utils.GetMessageText(proxyParams, m.cfg.Server.ChannelName)

		msg := tgbotapi.NewMessageToChannel(m.cfg.Server.ChannelName, text)
		msg.ParseMode = tgbotapi.ModeHTML
		msg.ReplyMarkup = buttons.CmdButtons(url)

		if _, err := m.bot.Send(msg); err != nil {
			m.logger.Error().Err(err).Msg("Error sending message")
			return
		}
		time.Sleep(time.Second)
	}
}

func NewMessageUseCase(bot *tgbotapi.BotAPI, cfg *config.Config, logger zerolog.Logger) messages.UseCase {
	return &messageUC{bot: bot, cfg: cfg, logger: logger}
}
