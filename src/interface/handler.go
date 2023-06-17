package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	domain "github.com/Keyl0ve/coffee-taste-app/src/domain/model"
	"github.com/Keyl0ve/coffee-taste-app/src/usecase"
)

type ServiceDriver struct {
	Controller usecase.Usecase
}

func NewServiceDriver(controller *usecase.ChatToolUsecase) *ServiceDriver {
	return &ServiceDriver{
		Controller: controller,
	}
}

// curl -H 'channelID:ChannelID_testChannel1_12345' http://localhost:8080/api/message/get/send
func (s *ServiceDriver) MessageGetSend(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// channelID を header から取得
	channelID := r.Header.Get("channelID")

	// channelID が空文字だったらエラーを返す
	if channelID == "" {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", "header の情報が不足しています。"))
	}

	messages, err := s.Controller.GetMessageByIsSend(ctx, domain.ChannelID(channelID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Messages []domain.Message `json:"messages"`
	}{
		Messages: messages,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl -H 'channelID:ChannelID_testChannel1_12345' -H 'userID:UserID_testUser1_12345' http://localhost:8080/api/message/get/notsend
func (s *ServiceDriver) MessageGetNotSend(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// channelID, UserID を header から取得
	channelID := r.Header.Get("channelID")
	userID := r.Header.Get("userID")

	// channelID, channelID が空文字だったらエラーを返す
	if channelID == "" || userID == "" {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", "header の情報が不足しています。"))
	}

	messages, err := s.Controller.GetMessageByNotIsSend(ctx, domain.ChannelID(channelID), domain.UserID(userID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Messages []domain.Message `json:"messages"`
	}{
		Messages: messages,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl http://localhost:8080/api/channel/get
func (s *ServiceDriver) ChannelGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	channels, err := s.Controller.GetChannels(ctx)
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Channels []domain.Channel `json:"channels"`
	}{
		Channels: channels,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl -H 'channelName:general' http://localhost:8080/api/channel/create
func (s *ServiceDriver) ChannelCreate(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	channelName := r.Header.Get("channelName")
	channel := domain.NewChannel(channelName, now)

	err := s.Controller.CreateChannel(ctx, channel)
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Success bool `json:"Success"`
	}{
		Success: true,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl -H 'userID:UserID_testUser1_12345' http://localhost:8080/api/user/get
func (s *ServiceDriver) UserGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	user, err := s.Controller.GetUserByUserID(ctx, domain.UserID(userID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		User domain.User `json:"User"`
	}{
		User: user,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl -H 'channelID:ChannelID_testChannel1_12345' -H 'userID:UserID_testUser1_12345' http://localhost:8080/api/join/delete
func (s *ServiceDriver) JoinDelete(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	channelID := r.Header.Get("channelID")

	// TODO: 既に脱退している場合、エラーを返す
	err := s.Controller.DeleteChannelByUserIDAndChannelID(ctx, domain.UserID(userID), domain.ChannelID(channelID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Success bool `json:"Success"`
	}{
		Success: true,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl -H 'channelID:ChannelID_testChannel1_12345' -H 'userID:UserID_testUser1_12345' http://localhost:8080/api/join/create
func (s *ServiceDriver) JoinCreate(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	channelID := r.Header.Get("channelID")
	now := time.Now()

	// userID から User 情報を取得
	user, err := s.Controller.GetUserByUserID(ctx, domain.UserID(userID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// channelID から Channel 情報を取得
	channel, err := s.Controller.GetChannelByChannelID(ctx, domain.ChannelID(channelID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	if userID == "" || user.UserName == "" || channelID == "" || channel.ChannelName == "" {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("ユーザー、もしくはチャンネルが不正です。"))
	}

	join := domain.NewJoinChannelToUser(domain.UserID(userID), user.UserName, domain.ChannelID(channelID), channel.ChannelName, now, now)

	// TODO: 既に参加している場合、エラーを返す
	err = s.Controller.CreateChannelConnection(ctx, join)
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Success bool `json:"Success"`
	}{
		Success: true,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl -H 'channelID:ChannelID_testChannel1_12345' http://localhost:8080/api/join/get/user
func (s *ServiceDriver) JoinGetUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	channelID := r.Header.Get("channelID")
	if channelID == "" {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", "header の情報が不足しています。"))
	}

	joinChannelToUsers, err := s.Controller.GetUserByChannelID(ctx, domain.ChannelID(channelID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Join []domain.JoinChannelToUser `json:"join"`
	}{
		Join: joinChannelToUsers,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl -H 'userID:UserID_testUser1_12345' http://localhost:8080/api/join/get/channel
func (s *ServiceDriver) JoinGetChannel(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// UserID を header から取得
	userID := r.Header.Get("userID")

	// userID が空文字だったらエラーを返す
	if userID == "" {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", "header の情報が不足しています。"))
	}

	joinChannelToUsers, err := s.Controller.GetChannelByUserID(ctx, domain.UserID(userID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Join []domain.JoinChannelToUser `json:"join"`
	}{
		Join: joinChannelToUsers,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}
