package usecase

import (
	"context"

	domain "github.com/Keyl0ve/coffee-taste-app/src/domain/model"
	"github.com/Keyl0ve/coffee-taste-app/src/domain/repository"
)

type Usecase interface {

	// メッセージに関する UseCase

	// ユーザーは、【チャット】の【所属メンバー】に【メッセージ】を送信する
	CreateMessage(ctx context.Context, message *domain.Message) error
	GetMessageByIsSend(ctx context.Context, channelID domain.ChannelID) ([]domain.Message, error)
	// ユーザーは、【送信予定のメッセージ】を全て確認できる
	GetMessageByNotIsSend(ctx context.Context, channelID domain.ChannelID, userID domain.UserID) ([]domain.Message, error)
	// ユーザーは、【送信予定のメッセージ】を編集・削除できる
	UpdateMessageByNotIsSend(ctx context.Context, afterMessage *domain.Message) error

	// ユーザーに関する UseCase

	// User を作成する
	CreateUser(ctx context.Context, user *domain.User) error
	// ChannelID でユーザーを取得する
	GetUserByChannelID(ctx context.Context, channelID domain.ChannelID) ([]domain.JoinChannelToUser, error)
	// 指定したユーザーの情報を取得する
	GetUserByUserID(ctx context.Context, userID domain.UserID) (domain.User, error)

	// チャンネルに関する UseCase

	// チャンネルを作成する
	CreateChannel(ctx context.Context, channel *domain.Channel) error
	// ユーザーは、【チャット】の【名前】,【所属メンバー】などを設定する。
	EditChannelConfig(ctx context.Context, beforeChannel *domain.Channel, afterChannel *domain.Channel) error
	// 全チャンネルを取得する
	GetChannels(ctx context.Context) ([]domain.Channel, error)
	// ChannelID でチャンネルを取得する
	GetChannelByChannelID(ctx context.Context, channelID domain.ChannelID) (domain.Channel, error)
	// UserID でチャンネルを取得する
	GetChannelByUserID(ctx context.Context, userID domain.UserID) ([]domain.JoinChannelToUser, error)
	// 指定したチャンネルの指定したユーザーを脱退させる
	DeleteChannelByUserIDAndChannelID(ctx context.Context, userID domain.UserID, channelID domain.ChannelID) error
	// 指定したチャンネルにユーザーを参加させる
	CreateChannelConnection(ctx context.Context, joinChannelToUser *domain.JoinChannelToUser) error
	// チャンネルを削除する
	DeleteChannel(ctx context.Context, channelD domain.ChannelID) error
}

type ChatToolUsecase struct {
	UserRepo    repository.UserRepository
	ChannelRepo repository.ChannelRepository
	MessageRepo repository.MessageRepository
	JoinRepo    repository.JoinChannelToUserRepository
}

func NewChatToolUsecase(userRepo repository.UserRepository, channelRepo repository.ChannelRepository, messageRepo repository.MessageRepository, joinRepo repository.JoinChannelToUserRepository) *ChatToolUsecase {
	return &ChatToolUsecase{UserRepo: userRepo, ChannelRepo: channelRepo, MessageRepo: messageRepo, JoinRepo: joinRepo}
}

func (c *ChatToolUsecase) CreateMessage(ctx context.Context, message *domain.Message) error {
	err := c.MessageRepo.CreateMessage(ctx, message)
	return err
}

func (c *ChatToolUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	err := c.UserRepo.CreateUser(ctx, user)
	return err
}

func (c *ChatToolUsecase) GetUserByChannelID(ctx context.Context, channelID domain.ChannelID) ([]domain.JoinChannelToUser, error) {
	join, err := c.JoinRepo.GetJoinByChannelID(ctx, channelID)
	return join, err
}

func (c *ChatToolUsecase) GetUserByUserID(ctx context.Context, userID domain.UserID) (domain.User, error) {
	user, err := c.UserRepo.GetUser(ctx, userID)
	return user, err
}

func (c *ChatToolUsecase) CreateChannel(ctx context.Context, channel *domain.Channel) error {
	err := c.ChannelRepo.CreateChannel(ctx, channel)
	return err
}

func (c *ChatToolUsecase) EditChannelConfig(ctx context.Context, beforeChannel *domain.Channel, afterChannel *domain.Channel) error {
	err := c.ChannelRepo.UpdateChannel(ctx, "", afterChannel)
	return err
}

func (c *ChatToolUsecase) GetChannelByChannelID(ctx context.Context, channelID domain.ChannelID) (domain.Channel, error) {
	channel, err := c.ChannelRepo.GetChannelByChannelID(ctx, channelID)
	return channel, err
}

func (c *ChatToolUsecase) DeleteChannelByUserIDAndChannelID(ctx context.Context, userID domain.UserID, channelID domain.ChannelID) error {
	err := c.JoinRepo.DeleteConnectionUserIDToChannelID(ctx, userID, channelID)
	return err
}

func (c *ChatToolUsecase) CreateChannelConnection(ctx context.Context, joinChannelToUser *domain.JoinChannelToUser) error {
	err := c.JoinRepo.CreateConnectionUserIDToChannelID(ctx, joinChannelToUser)
	return err
}

func (c *ChatToolUsecase) DeleteChannel(ctx context.Context, channelD domain.ChannelID) error {
	err := c.ChannelRepo.DeleteChannel(ctx, channelD)
	return err
}

func (c *ChatToolUsecase) GetMessageByIsSend(ctx context.Context, channelID domain.ChannelID) ([]domain.Message, error) {
	messages, err := c.MessageRepo.GetAllSendMessages(ctx, channelID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (c *ChatToolUsecase) GetMessageByNotIsSend(ctx context.Context, channelID domain.ChannelID, userID domain.UserID) ([]domain.Message, error) {
	messages, err := c.MessageRepo.GetMessagesByChannelIDAndIsNotSendAndUserID(ctx, channelID, userID)
	if err != nil {
		return nil, err
	}
	return messages, nil

}

func (c *ChatToolUsecase) GetChannels(ctx context.Context) ([]domain.Channel, error) {
	channel, err := c.ChannelRepo.GetChannels(ctx)
	if err != nil {
		return nil, err
	}
	return channel, nil
}

func (c *ChatToolUsecase) GetChannelByUserID(ctx context.Context, userID domain.UserID) ([]domain.JoinChannelToUser, error) {
	joinChannelToUsers, err := c.JoinRepo.GetJoinByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return joinChannelToUsers, nil
}

func (c *ChatToolUsecase) UpdateMessageByNotIsSend(ctx context.Context, afterMessage *domain.Message) error {
	err := c.MessageRepo.UpdateMessage(ctx, afterMessage)
	return err
}
