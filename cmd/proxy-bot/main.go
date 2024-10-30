package main

import (
	"github.com/RajabovIlyas/proxy-bot/internal/pkg/app"
	"github.com/RajabovIlyas/proxy-bot/internal/pkg/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	appLogger, err := logger.InitLogger()

	if err != nil {
		log.Fatal().Err(err).Msg("Error initializing logger")
		return
	}

	err = app.Run(appLogger)

	if err != nil {
		appLogger.Fatal().Msg(err.Error())
	}
}
