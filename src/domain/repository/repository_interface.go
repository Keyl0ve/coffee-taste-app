package repository

import (
	"context"

	domain "github.com/Keyl0ve/coffee-taste-app/src/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUsers(ctx context.Context) ([]domain.User, error)
}

type CoffeeRepository interface {
	CreateCoffee(ctx context.Context, coffee *domain.Coffee) error
	GetCoffees(ctx context.Context) ([]domain.Coffee, error)
}
