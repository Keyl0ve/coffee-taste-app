package infra

import (
	"context"
	"database/sql"
	"log"

	domain "github.com/Keyl0ve/coffee-taste-app/src/domain/model"
)

type JoinCoffeeToUserRepository struct {
	Conn *sql.DB
}

func NewJoinCoffeeToUserRepository(conn *sql.DB) *JoinCoffeeToUserRepository {
	return &JoinCoffeeToUserRepository{Conn: conn}
}

func ScanJoinCoffeeToUsers(rows *sql.Rows) ([]domain.JoinCoffeeToUser, int, error) {
	joinCoffeeToUsers := make([]domain.JoinCoffeeToUser, 0)

	for rows.Next() {
		var v domain.JoinCoffeeToUser
		if err := rows.Scan(&v.UserID, &v.UserName, &v.CoffeeID, &v.CoffeeName); err != nil {
			log.Printf("[ERROR] scan ScanJoinCoffeeToUsers: %+v", err)
			return nil, 0, err
		}
		joinCoffeeToUsers = append(joinCoffeeToUsers, v)
	}

	return joinCoffeeToUsers, len(joinCoffeeToUsers), nil
}

func (j JoinCoffeeToUserRepository) GetJoinByUserID(ctx context.Context, userID domain.UserID) ([]domain.JoinCoffeeToUser, error) {
	query := "SELECT * FROM joinCoffeeToUser WHERE user_id = ?"
	rows, err := j.Conn.QueryContext(ctx, query, userID)
	if err != nil {
		log.Printf("[ERROR] can't get GetCoffeeIDsByUserID: %+v", err)
		return nil, err
	}

	joinCoffeeToUsers, _, err := ScanJoinCoffeeToUsers(rows)
	if err != nil {
		log.Printf("[ERROR] can not scan Coffees: %+v", err)
		return nil, err
	}

	return joinCoffeeToUsers, nil
}

func (j JoinCoffeeToUserRepository) GetJoinByCoffeeID(ctx context.Context, coffeeID domain.CoffeeID) ([]domain.JoinCoffeeToUser, error) {
	query := "SELECT * FROM joinCoffeeToUser WHERE coffee_id = ?"
	rows, err := j.Conn.QueryContext(ctx, query, coffeeID)
	if err != nil {
		log.Printf("[ERROR] can't get GetUserIDsByCoffeeID: %+v", err)
		return nil, err
	}

	joinCoffeeToUsers, _, err := ScanJoinCoffeeToUsers(rows)
	if err != nil {
		log.Printf("[ERROR] can not scan Coffees: %+v", err)
		return nil, err
	}

	return joinCoffeeToUsers, nil
}

func (j JoinCoffeeToUserRepository) CreateConnectionUserIDToCoffeeID(ctx context.Context, join *domain.JoinCoffeeToUser) error {
	query := "INSERT INTO joinCoffeeToUser (user_ID, user_name, coffee_ID, coffee_name, created_at, updated_at) VALUES (?,?,?,?) "
	_, err := j.Conn.ExecContext(ctx, query, join.UserID, join.UserName, join.CoffeeID, join.CoffeeName)
	if err != nil {
		log.Printf("[ERROR] can't create CreateConnectionUserIDToCoffeeID: %+v", err)
		return nil
	}

	return nil
}

func (j JoinCoffeeToUserRepository) DeleteConnectionUserIDToCoffeeID(ctx context.Context, userid domain.UserID, coffeeID domain.CoffeeID) error {
	query := "DELETE FROM joinCoffeeToUser WHERE user_id = ? AND coffee_id = ?"
	_, err := j.Conn.ExecContext(ctx, query, userid, coffeeID)
	if err != nil {
		log.Printf("[ERROR] can't delete DeleteConnectionUserIDToCoffeeID: %+v", err)
		return nil
	}

	return nil
}
