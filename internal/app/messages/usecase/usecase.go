package usecase

import (
	"errors"
	"github.com/RajabovIlyas/proxy-bot/config"
	"github.com/RajabovIlyas/proxy-bot/internal/app/messages"
	"github.com/RajabovIlyas/proxy-bot/internal/app/utils"
	"github.com/RajabovIlyas/proxy-bot/internal/pkg/buttons"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
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

	if update.Message == nil {
		m.SendErrorMessage(update, errors.New("incorrect input"))
		return
	}

	proxy := update.Message.Text
	proxyParams, err := utils.GetProxyUrlParams(proxy)

	if err != nil {
		m.SendErrorMessage(update, err)
		return
	}

	text := utils.GetMessageText(proxyParams, m.cfg.Server.ChannelName)

	msg := tgbotapi.NewMessageToChannel(m.cfg.Server.ChannelName, text)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = buttons.CmdButtons(proxy)

	if _, err := m.bot.Send(msg); err != nil {
		panic(err)
	}
}

func NewMessageUseCase(bot *tgbotapi.BotAPI, cfg *config.Config, logger zerolog.Logger) messages.UseCase {
	return &messageUC{bot: bot, cfg: cfg, logger: logger}
}
