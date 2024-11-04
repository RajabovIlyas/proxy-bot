package server

import (
	"context"
	"github.com/RajabovIlyas/proxy-bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	messageHandlers, err := s.mapHandlers()

	if err != nil {
		return err
	}

	go func() {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates := s.bot.GetUpdatesChan(u)
		// Loop through each update.
		for update := range updates {
			messageHandlers.Messages(update)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	_, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)

	defer shutdown()

	s.logger.Info().Msg("Server Exited Properly")
	return nil
}
