package repository

import (
	"go.uber.org/zap"
	"wanna-go-out-api/domain"
	"wanna-go-out-api/infrastructure"
)

type messageRepository struct {
	Orm    *infrastructure.Orm
	Logger *zap.Logger
}

func (mr *messageRepository) Save(message *domain.Message) (*domain.Message, error) {

	result := mr.Orm.Db.Create(message)

	if result.Error != nil {
		return nil, result.Error
	}

	return message, nil

}

func (mr *messageRepository) Get() (messages []domain.Message) {

	mr.Orm.Db.Find(&messages)

	return messages
}

func NewMessageRepository(orm *infrastructure.Orm, logger *zap.Logger) domain.MessageRepository {
	return &messageRepository{
		Orm:    orm,
		Logger: logger,
	}

}
