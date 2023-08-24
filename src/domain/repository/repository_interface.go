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

type CoffeeRepository interface {
	CreateCoffee(ctx context.Context, coffee *domain.Coffee) error
	GetCoffees(ctx context.Context) ([]domain.Coffee, error)
	GetCoffeeByCoffeeID(ctx context.Context, coffeeID domain.CoffeeID) (domain.Coffee, error)
	UpdateCoffee(ctx context.Context, coffeeID domain.CoffeeID, updatedCoffee *domain.Coffee) error
	DeleteCoffee(ctx context.Context, coffeeID domain.CoffeeID) error
}

type JoinCoffeeToUserRepository interface {
	// userID を指定して coffeeID と coffeeName を GET
	GetJoinByUserID(ctx context.Context, userID domain.UserID) ([]domain.JoinCoffeeToUser, error)
	// coffeeID を指定して userID と userName を GET
	GetJoinByCoffeeID(ctx context.Context, coffeeID domain.CoffeeID) ([]domain.JoinCoffeeToUser, error)
	// チャンネルに入会したときに実行される
	CreateConnectionUserIDToCoffeeID(ctx context.Context, join *domain.JoinCoffeeToUser) error
	// チャンネルから脱退したときに実行される
	DeleteConnectionUserIDToCoffeeID(ctx context.Context, userid domain.UserID, coffeeID domain.CoffeeID) error
}
