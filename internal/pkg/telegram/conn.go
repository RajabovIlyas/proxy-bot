package telegram

import (
	"github.com/RajabovIlyas/proxy-bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
)

func InitBot(cfgServer config.ServerConfig, logger zerolog.Logger) *tgbotapi.BotAPI {

	tgbotapi.SetLogger(&logger)

	bot, err := tgbotapi.NewBotAPI(cfgServer.TelegramApiToken)
	if err != nil {
		logger.Panic().Err(err).Msg("failed to init telegram bot")
	}
	bot.Debug = true
	return bot
}
