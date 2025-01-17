package adapters

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wanna-go-out-api/domain"
)

type MessageController struct {
	Message domain.MessageUseCase
}

func (mc *MessageController) GetMessagesHandler(c *gin.Context) {
	var messages []domain.Message

	messages = mc.Message.GetAllMessage()

	c.JSON(http.StatusOK, messages)

}

func (mc *MessageController) CreateMessageHandler(c *gin.Context) {
	var message domain.Message

	err := c.ShouldBind(&message)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	insertedMessage, err := mc.Message.SaveMessage(&message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, insertedMessage)

}
