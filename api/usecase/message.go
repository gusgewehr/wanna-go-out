package usecase

import (
	"go.uber.org/zap"
	"wanna-go-out-api/domain"
)

type messageUseCase struct {
	messageRepository domain.MessageRepository
	logger            *zap.Logger
}

func (mu *messageUseCase) SaveMessage(message *domain.Message) (*domain.Message, error) {

	insertedMessage, err := mu.messageRepository.Save(message)
	if err != nil {
		return nil, err
	}

	return insertedMessage, nil

}

func (mu *messageUseCase) GetAllMessage() (messages []domain.Message) {

	messages = mu.messageRepository.Get()

	return messages
}

func NewMessageUseCase(logger *zap.Logger, messageRepository domain.MessageRepository) domain.MessageUseCase {
	return &messageUseCase{
		messageRepository: messageRepository,
		logger:            logger,
	}
}
