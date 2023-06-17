package infra

import (
	"context"
	"database/sql"
	"log"

	domain "github.com/Keyl0ve/coffee-taste-app/src/domain/model"
)

type JoinChannelToUserRepository struct {
	Conn *sql.DB
}

func NewJoinChannelToUserRepository(conn *sql.DB) *JoinChannelToUserRepository {
	return &JoinChannelToUserRepository{Conn: conn}
}

func ScanJoinChannelToUsers(rows *sql.Rows) ([]domain.JoinChannelToUser, int, error) {
	joinChannelToUsers := make([]domain.JoinChannelToUser, 0)

	for rows.Next() {
		var v domain.JoinChannelToUser
		if err := rows.Scan(&v.UserID, &v.UserName, &v.ChannelID, &v.ChannelName, &v.CreatedAt, &v.UpdatedAt); err != nil {
			log.Printf("[ERROR] scan ScanJoinChannelToUsers: %+v", err)
			return nil, 0, err
		}
		joinChannelToUsers = append(joinChannelToUsers, v)
	}

	return joinChannelToUsers, len(joinChannelToUsers), nil
}

func (j JoinChannelToUserRepository) GetJoinByUserID(ctx context.Context, userID domain.UserID) ([]domain.JoinChannelToUser, error) {
	query := "SELECT * FROM joinChannelToUser WHERE user_id = ?"
	rows, err := j.Conn.QueryContext(ctx, query, userID)
	if err != nil {
		log.Printf("[ERROR] can't get GetChannelIDsByUserID: %+v", err)
		return nil, err
	}

	joinChannelToUsers, _, err := ScanJoinChannelToUsers(rows)
	if err != nil {
		log.Printf("[ERROR] can not scan Channels: %+v", err)
		return nil, err
	}

	return joinChannelToUsers, nil
}

func (j JoinChannelToUserRepository) GetJoinByChannelID(ctx context.Context, channelID domain.ChannelID) ([]domain.JoinChannelToUser, error) {
	query := "SELECT * FROM joinChannelToUser WHERE channel_id = ?"
	rows, err := j.Conn.QueryContext(ctx, query, channelID)
	if err != nil {
		log.Printf("[ERROR] can't get GetUserIDsByChannelID: %+v", err)
		return nil, err
	}

	joinChannelToUsers, _, err := ScanJoinChannelToUsers(rows)
	if err != nil {
		log.Printf("[ERROR] can not scan Channels: %+v", err)
		return nil, err
	}

	return joinChannelToUsers, nil
}

func (j JoinChannelToUserRepository) CreateConnectionUserIDToChannelID(ctx context.Context, join *domain.JoinChannelToUser) error {
	query := "INSERT INTO joinChannelToUser (user_ID, user_name, channel_ID, channel_name, created_at, updated_at) VALUES (?,?,?,?,?,?) "
	_, err := j.Conn.ExecContext(ctx, query, join.UserID, join.UserName, join.ChannelID, join.ChannelName, join.CreatedAt, join.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR] can't create CreateConnectionUserIDToChannelID: %+v", err)
		return nil
	}

	return nil
}

func (j JoinChannelToUserRepository) DeleteConnectionUserIDToChannelID(ctx context.Context, userid domain.UserID, channelID domain.ChannelID) error {
	query := "DELETE FROM joinChannelToUser WHERE user_id = ? AND channel_id = ?"
	_, err := j.Conn.ExecContext(ctx, query, userid, channelID)
	if err != nil {
		log.Printf("[ERROR] can't delete DeleteConnectionUserIDToChannelID: %+v", err)
		return nil
	}

	return nil
}
