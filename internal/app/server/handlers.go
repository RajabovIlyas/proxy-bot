package server

import (
	messageBotHandler "github.com/RajabovIlyas/proxy-bot/internal/app/messages/delivery"
	messageUseCase "github.com/RajabovIlyas/proxy-bot/internal/app/messages/usecase"
)

func (s *Server) mapHandlers() error {

	// Init useCases
	messageUC := messageUseCase.NewMessageUseCase(s.bot, s.cfg, s.logger)

	// Init handlers
	messageHandlers := messageBotHandler.NewMessageHandlers(messageUC, s.bot, s.cfg, s.logger)

	messageHandlers.InitMessageTGBot()

	return nil
}
