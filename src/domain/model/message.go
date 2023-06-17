package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MessageID string

func NewMessageID(messageTitle string) MessageID {
	uuid, _ := uuid.NewRandom()
	return MessageID(fmt.Sprintf("MessageID_%s_%s", messageTitle, uuid))
}

// func (m *Message) ValidateMessage() error {
// 	if m.MessageID == "" {
// 		return fmt.Errorf("messageID is empty")
// 	}
// 	if m.MessageBody == "" {
// 		return fmt.Errorf("messageBody is empty")
// 	}
// 	if m.Author == "" {
// 		return fmt.Errorf("author is empty")
// 	}
// 	if m.ChannelID == "" {
// 		return fmt.Errorf("channelID is empty")
// 	}
// 	return nil
// }

// func (ms Messages) ValidateMessages() error {
// 	if len(ms) > 100 {
// 		return fmt.Errorf("too many messages")
// 	}
// 	return nil
// }

type Messages []*Message

type Message struct {
	MessageID   MessageID
	MessageBody string
	Author      UserID
	ChannelID   ChannelID
	IsSend      bool
	SendAt      time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
