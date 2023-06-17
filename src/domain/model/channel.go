package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ChannelID string

func NewChannelID(name string) ChannelID {
	uuid, _ := uuid.NewRandom()
	return ChannelID(fmt.Sprintf("ChannelID_%s_%s", name, uuid))
}

type Channel struct {
	ChannelID   ChannelID
	ChannelName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewChannel(name string, now time.Time) *Channel {
	return &Channel{
		ChannelID:   NewChannelID(name),
		ChannelName: name,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
