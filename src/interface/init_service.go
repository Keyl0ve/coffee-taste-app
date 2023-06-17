package handler

import (
	"database/sql"

	"github.com/Keyl0ve/coffee-taste-app/src/infra"
	"github.com/Keyl0ve/coffee-taste-app/src/usecase"
)

func InitService(conn *sql.DB) Service {
	userRepo := infra.NewUserRepository(conn)
	channelRepo := infra.NewChannelRepository(conn)
	messageRepo := infra.NewMessageRepository(conn)
	joinRepo := infra.NewJoinChannelToUserRepository(conn)

	usecaseInterface := usecase.NewChatToolUsecase(userRepo, channelRepo, messageRepo, joinRepo)
	service := NewServiceDriver(usecaseInterface)
	return service
}
