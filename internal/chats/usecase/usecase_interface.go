package usecase

import (
	"context"
	"net/http"

	chatModel "github.com/go-park-mail-ru/2024_2_EaglesDesigner/internal/chats/models"
	"github.com/google/uuid"
)

type ChatUsecase interface {
	GetChats(ctx context.Context, cookie []*http.Cookie) ([]chatModel.ChatDTOOutput, error)
	AddUsersIntoChat(ctx context.Context, cookie []*http.Cookie, user_ids []uuid.UUID, chat_id uuid.UUID) (chatModel.AddedUsersIntoChatDTO, error)

	// CanUserWriteInChat проверяет может ли юзер писать в чат
	AddNewChat(ctx context.Context, cookie []*http.Cookie, chat chatModel.ChatDTOInput) (chatModel.ChatDTOOutput, error)

	DeleteChat(ctx context.Context, chatId uuid.UUID, userId uuid.UUID) error
	UpdateChat(ctx context.Context, chatId uuid.UUID, chatUpdate chatModel.ChatUpdate, userId uuid.UUID) (chatModel.ChatUpdateOutput, error)

	DeleteUsersFromChat(ctx context.Context, userID uuid.UUID, chatId uuid.UUID, usertToDelete chatModel.DeleteUsersFromChatDTO) (chatModel.DeletdeUsersFromChatDTO, error)
	GetUsersFromChat(ctx context.Context, chatId uuid.UUID, userId uuid.UUID) (chatModel.UsersInChatDTO, error)
}
