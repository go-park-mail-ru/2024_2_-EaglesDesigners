package usecase

import (
	"context"
	"net/http"

	chatModel "github.com/go-park-mail-ru/2024_2_EaglesDesigner/main_app/internal/chats/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source=usecase_interface.go -destination=mocks/mocks.go

type ChatUsecase interface {
	GetChats(ctx context.Context, cookie []*http.Cookie) ([]chatModel.ChatDTOOutput, error)
	AddUsersIntoChatWithCheckPermission(ctx context.Context, user_ids []uuid.UUID, chat_id uuid.UUID) (chatModel.AddedUsersIntoChatDTO, error)

	// CanUserWriteInChat проверяет может ли юзер писать в чат
	AddNewChat(ctx context.Context, cookie []*http.Cookie, chat chatModel.ChatDTOInput) (chatModel.ChatDTOOutput, error)

	DeleteChat(ctx context.Context, chatId uuid.UUID, userId uuid.UUID) error
	UpdateChat(ctx context.Context, chatId uuid.UUID, chatUpdate chatModel.ChatUpdate, userId uuid.UUID) (chatModel.ChatUpdateOutput, error)

	DeleteUsersFromChat(ctx context.Context, userID uuid.UUID, chatId uuid.UUID, usertToDelete chatModel.DeleteUsersFromChatDTO) (chatModel.DeletdeUsersFromChatDTO, error)

	JoinChannel(ctx context.Context, userId uuid.UUID, channelId uuid.UUID) error
	
	// UserLeaveChat удаляет владельца обращения из чата
	UserLeaveChat(ctx context.Context, userId uuid.UUID, chatId uuid.UUID) error

	GetChatInfo(ctx context.Context, chatId uuid.UUID, userId uuid.UUID) (chatModel.ChatInfoDTO, error)

	AddBranch(ctx context.Context, chatId uuid.UUID, messageID uuid.UUID, userId uuid.UUID) (chatModel.AddBranch, error)

	SearchChats(ctx context.Context, userID uuid.UUID, keyWord string) (chatModel.SearchChatsDTO, error)

	// grpc
	GetUserChats(ctx context.Context, userId string) (chatIds []string, err error)
	GetUsersFromChat(ctx context.Context, chatId string) (userIds []string, err error)
}
