package infra

import (
	"context"
	"database/sql"
	"log"

	domain "github.com/Keyl0ve/coffee-taste-app/src/domain/model"
)

type UserRepository struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{Conn: conn}
}

func (u UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	// 新規ユーザーの登録
	query := "INSERT INTO user (user_ID, user_name, password, created_at, updated_at) VALUES (?,?,?,?,?) "
	_, err := u.Conn.ExecContext(ctx, query, user.UserID, user.UserName, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR] can't create User: %+v", err)
		return nil
	}

	return nil
}
func (u UserRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	query := "SELECT * FROM user"
	rows, err := u.Conn.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[ERROR] can't get Users: %+v", err)
		return []domain.User{}, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.UserID, &user.UserName, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.Printf("[ERROR] scan ScanUser: %+v", err)
			return []domain.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}
