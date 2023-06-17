package infra

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	domain "github.com/Keyl0ve/coffee-taste-app/src/domain/model"
)

type MessageRepository struct {
	Conn *sql.DB
}

func NewMessageRepository(conn *sql.DB) *MessageRepository {
	return &MessageRepository{Conn: conn}
}

func ScanMessages(rows *sql.Rows) ([]domain.Message, int, error) {
	messages := make([]domain.Message, 0)
	for rows.Next() {
		var v domain.Message
		if err := rows.Scan(&v.MessageID, &v.MessageBody, &v.Author, &v.ChannelID, &v.IsSend, &v.SendAt, &v.CreatedAt, &v.UpdatedAt); err != nil {
			log.Printf("[ERROR] scan ScanChannels: %+v", err)
			return nil, 0, err
		}
		fmt.Sprintf("message: %+v/n", v)
		fmt.Println(v.MessageBody)
		messages = append(messages, v)
	}

	fmt.Println(messages, len(messages))
	return messages, len(messages), nil
}

func (m *MessageRepository) CreateMessage(ctx context.Context, message *domain.Message) error {
	query := "INSERT INTO message (message_id, message_body, author, channel_id, is_send, send_at, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?) "
	_, err := m.Conn.ExecContext(ctx, query, message.MessageID, message.MessageBody, message.Author, message.ChannelID, message.IsSend, message.SendAt, message.CreatedAt, message.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR] can't create Message: %+v", err)
		return nil
	}

	return nil
}

func (m *MessageRepository) GetAllSendMessages(ctx context.Context, channelID domain.ChannelID) ([]domain.Message, error) {
	query := "SELECT * FROM message WHERE channel_id = ? AND is_send = true"
	rows, err := m.Conn.QueryContext(ctx, query, channelID)
	if err != nil {
		log.Printf("[ERROR] can't get GetAllSendMessage: %+v", err)
		return nil, err
	}

	messages, _, err := ScanMessages(rows)
	if err != nil {
		log.Printf("[ERROR] can not scan GetAllSendMessages: %+v", err)
		return nil, err
	}

	fmt.Printf("message: %+v", &messages)

	return messages, nil
}

func (m *MessageRepository) GetMessagesByChannelIDAndIsNotSendAndUserID(ctx context.Context, channelID domain.ChannelID, userID domain.UserID) ([]domain.Message, error) {
	query := "SELECT * FROM message WHERE channel_id = ? AND is_send = false AND author = ?"
	rows, err := m.Conn.QueryContext(ctx, query, channelID, userID)
	if err != nil {
		log.Printf("[ERROR] can't get GetMessagesByChannelIDAndIsNotSendAndUserID: %+v", err)
		return nil, err
	}

	messages, _, err := ScanMessages(rows)
	if err != nil {
		log.Printf("[ERROR] can not scan GetMessagesByChannelIDAndIsNotSendAndUserID: %+v", err)
		return nil, err
	}

	return messages, nil
}

func (m *MessageRepository) UpdateMessage(ctx context.Context, updatedMessage *domain.Message) error {
	query := "UPDATE message set message_body = ? , is_send = ?, send_at = ?, updated_at = ? WHERE channel_id = ? "
	_, err := m.Conn.ExecContext(ctx, query, updatedMessage.MessageBody, updatedMessage.IsSend, updatedMessage.SendAt, updatedMessage.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR] can't UpdateMessage: %+v", err)
		return nil
	}

	return nil
}

func (m *MessageRepository) DeleteMessage(ctx context.Context, messageID domain.MessageID) error {
	query := "DELETE FROM message WHERE message_id = ?"
	_, err := m.Conn.ExecContext(ctx, query, messageID)
	if err != nil {
		log.Printf("[ERROR] can't DeleteMessage: %+v", err)
		return nil
	}

	return nil
}
