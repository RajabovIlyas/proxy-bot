package usecase

import (
	"errors"
	"github.com/RajabovIlyas/proxy-bot/config"
	"github.com/RajabovIlyas/proxy-bot/internal/app/proxy"
	"github.com/RajabovIlyas/proxy-bot/internal/app/utils"
	"github.com/RajabovIlyas/proxy-bot/internal/pkg/buttons"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
)

type proxyUC struct {
	cfg    *config.Config
	logger zerolog.Logger
}

func (p proxyUC) GetMessages(message *tgbotapi.Message) ([]tgbotapi.MessageConfig, error) {

	if message == nil {
		return nil, errors.New("incorrect input")
	}

	urls := utils.FilterURLs(message)

	if len(urls) == 0 {
		return nil, errors.New("incorrect input")
	}

	sendMessages := make([]tgbotapi.MessageConfig, len(urls))

	for _, url := range urls {

		proxyParams, err := utils.GetURLParams(url)

		if err != nil {
			continue
		}

		text := utils.GetMessageText(proxyParams, p.cfg.Server.ChannelName)

		msg := tgbotapi.NewMessageToChannel(p.cfg.Server.ChannelName, text)
		msg.ParseMode = tgbotapi.ModeHTML
		msg.ReplyMarkup = buttons.CmdButtons(url)

		sendMessages = append(sendMessages, msg)
	}

	return sendMessages, nil
}

func NewProxyUseCase(cfg *config.Config, logger zerolog.Logger) proxy.UseCase {
	return &proxyUC{cfg: cfg, logger: logger}
}
