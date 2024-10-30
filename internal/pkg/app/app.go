package app

import (
	"github.com/RajabovIlyas/proxy-bot/config"
	"github.com/RajabovIlyas/proxy-bot/internal/app/constants"
	"github.com/RajabovIlyas/proxy-bot/internal/app/server"
	"github.com/RajabovIlyas/proxy-bot/internal/pkg/telegram"
	"github.com/rs/zerolog"
)

func Run(logger zerolog.Logger) error {
	loadConfig, err := config.LoadConfig(constants.CONFIG_FILE_PATH)

	if err != nil {
		logger.Fatal().Err(err).Msg("app.Run: load config error:")
		return err
	}

	cfg, _ := config.ParseConfig(loadConfig, logger)

	bot := telegram.InitBot(cfg.Server, logger)

	s := server.NewServer(bot, cfg, logger)

	return s.Run()
}
