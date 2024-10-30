package server

import (
	"github.com/RajabovIlyas/proxy-bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
)

const (
	ctxTimeout = 5
)

type Server struct {
	bot    *tgbotapi.BotAPI
	cfg    *config.Config
	logger zerolog.Logger
}

func NewServer(bot *tgbotapi.BotAPI, cfg *config.Config, logger zerolog.Logger) *Server {
	return &Server{bot: bot, cfg: cfg, logger: logger}
}

func (s *Server) Run() error {

	if err := s.mapHandlers(); err != nil {
		return err
	}

	s.logger.Info().Msg("Server Exited Properly")
	return nil
}
