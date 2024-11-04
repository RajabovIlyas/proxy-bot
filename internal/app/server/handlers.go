package server

import (
	"github.com/RajabovIlyas/proxy-bot/internal/app/messages"
	messageBotHandler "github.com/RajabovIlyas/proxy-bot/internal/app/messages/delivery"
	messageUseCase "github.com/RajabovIlyas/proxy-bot/internal/app/messages/usecase"
	proxyUseCase "github.com/RajabovIlyas/proxy-bot/internal/app/proxy/usecase"
)

func (s *Server) mapHandlers() (messages.Handlers, error) {

	// Init useCases
	proxyUC := proxyUseCase.NewProxyUseCase(s.cfg, s.logger)
	messageUC := messageUseCase.NewMessageUseCase(s.bot, s.cfg, s.logger, proxyUC)

	// Init handlers
	messageHandlers := messageBotHandler.NewMessageHandlers(messageUC, s.bot, s.cfg, s.logger)

	return messageHandlers, nil
}
