package chatlist

import (
	"net/http"

	userModel "github.com/go-park-mail-ru/2024_2_EaglesDesigner/internal/auth/repository"
	chatModel "github.com/go-park-mail-ru/2024_2_EaglesDesigner/internal/chat_list/models"
)

type ChatRepository interface {
	GetUserChats(user *userModel.User) []chatModel.Chat
}

type ChatService interface {
	GetChats(cookie []*http.Cookie) ([]chatModel.Chat, error)
}
