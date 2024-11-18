package repository

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_EaglesDesigner/internal/messages/models"
	"github.com/google/uuid"
)

type MessageRepository interface {
	AddMessage(message models.Message, chatId uuid.UUID) error

	DeleteMessage(ctx context.Context, messageId uuid.UUID) error

	UpdateMessage(ctx context.Context, messageId uuid.UUID, newText string) error

	GetFirstMessages(ctx context.Context, chatId uuid.UUID) ([]models.Message, error)
	GetMessageById(ctx context.Context, messageId uuid.UUID) (models.Message, error)
	GetLastMessage(chatId uuid.UUID) (models.Message, error)
	GetAllMessagesAfter(ctx context.Context, chatId uuid.UUID, lastMessageId uuid.UUID) ([]models.Message, error)
}
