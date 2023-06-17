package model

import "time"

type JoinChannelToUser struct {
	UserID      UserID
	UserName    string
	ChannelID   ChannelID
	ChannelName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewJoinChannelToUser(userID UserID, userName string, channelID ChannelID, channelName string, createdAt time.Time, updatedAt time.Time) *JoinChannelToUser {
	return &JoinChannelToUser{
		UserID:      userID,
		UserName:    userName,
		ChannelID:   channelID,
		ChannelName: channelName,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
