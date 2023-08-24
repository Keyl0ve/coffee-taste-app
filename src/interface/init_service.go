package handler

import (
	"database/sql"

	"github.com/Keyl0ve/coffee-taste-app/src/infra"
	"github.com/Keyl0ve/coffee-taste-app/src/usecase"
)

func InitService(conn *sql.DB) Service {
	userRepo := infra.NewUserRepository(conn)
	coffeeRepo := infra.NewCoffeeRepository(conn)

	usecaseInterface := usecase.NewChatToolUsecase(userRepo, coffeeRepo)
	service := NewServiceDriver(usecaseInterface)
	return service
}
