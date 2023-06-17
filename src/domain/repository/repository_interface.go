package repository

import (
	"context"

	domain "github.com/Keyl0ve/coffee-taste-app/src/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, userID domain.UserID) (domain.User, error)
	UpdateUser(ctx context.Context, userID string, updatedUser *domain.User) error
	DeleteUser(ctx context.Context, userID domain.UserID) error
}

type ChannelRepository interface {
	CreateChannel(ctx context.Context, channel *domain.Channel) error
	GetChannels(ctx context.Context) ([]domain.Channel, error)
	GetChannelByChannelID(ctx context.Context, channelID domain.ChannelID) (domain.Channel, error)
	UpdateChannel(ctx context.Context, channelID domain.ChannelID, updatedChannel *domain.Channel) error
	DeleteChannel(ctx context.Context, channelID domain.ChannelID) error
}

type MessageRepository interface {
	// 新しいメッセージを作成する
	CreateMessage(ctx context.Context, message *domain.Message) error

	// ChannelID を指定して送信済みすべてのメッセージ一覧を取得
	// フロントエンドから 1 分に 1 回のリクエストを想定
	GetAllSendMessages(ctx context.Context, channelID domain.ChannelID) ([]domain.Message, error)

	// ChannelID を指定して特定の user の未送信のメッセージ一覧を取得
	GetMessagesByChannelIDAndIsNotSendAndUserID(ctx context.Context, channelID domain.ChannelID, userID domain.UserID) ([]domain.Message, error)

	// メッセージを更新する
	UpdateMessage(ctx context.Context, updatedMessage *domain.Message) error
	// メッセージを削除する
	DeleteMessage(ctx context.Context, messageID domain.MessageID) error
}

type JoinChannelToUserRepository interface {
	// userID を指定して channelID と channelName を GET
	GetJoinByUserID(ctx context.Context, userID domain.UserID) ([]domain.JoinChannelToUser, error)
	// channelID を指定して userID と userName を GET
	GetJoinByChannelID(ctx context.Context, channelID domain.ChannelID) ([]domain.JoinChannelToUser, error)
	// チャンネルに入会したときに実行される
	CreateConnectionUserIDToChannelID(ctx context.Context, join *domain.JoinChannelToUser) error
	// チャンネルから脱退したときに実行される
	DeleteConnectionUserIDToChannelID(ctx context.Context, userid domain.UserID, channelID domain.ChannelID) error
}
