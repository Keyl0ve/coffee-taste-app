package usecase

import (
	"context"

	domain "github.com/Keyl0ve/coffee-taste-app/src/domain/model"
	"github.com/Keyl0ve/coffee-taste-app/src/domain/repository"
)

type Usecase interface {

	// メッセージに関する UseCase

	// ユーザーに関する UseCase

	// User を作成する
	CreateUser(ctx context.Context, user *domain.User) error
	// CoffeeID でユーザーを取得する
	GetUserByCoffeeID(ctx context.Context, coffeeID domain.CoffeeID) ([]domain.JoinCoffeeToUser, error)
	// 指定したユーザーの情報を取得する
	GetUserByUserID(ctx context.Context, userID domain.UserID) (domain.User, error)

	// チャンネルに関する UseCase

	// チャンネルを作成する
	CreateCoffee(ctx context.Context, coffee *domain.Coffee) error
	// ユーザーは、【チャット】の【名前】,【所属メンバー】などを設定する。
	EditCoffeeConfig(ctx context.Context, beforeCoffee *domain.Coffee, afterCoffee *domain.Coffee) error
	// 全チャンネルを取得する
	GetCoffees(ctx context.Context) ([]domain.Coffee, error)
	// CoffeeID でチャンネルを取得する
	GetCoffeeByCoffeeID(ctx context.Context, coffeeID domain.CoffeeID) (domain.Coffee, error)
	// UserID でチャンネルを取得する
	GetCoffeeByUserID(ctx context.Context, userID domain.UserID) ([]domain.JoinCoffeeToUser, error)
	// 指定したチャンネルの指定したユーザーを脱退させる
	DeleteCoffeeByUserIDAndCoffeeID(ctx context.Context, userID domain.UserID, coffeeID domain.CoffeeID) error
	// 指定したチャンネルにユーザーを参加させる
	CreateCoffeeConnection(ctx context.Context, joinCoffeeToUser *domain.JoinCoffeeToUser) error
	// チャンネルを削除する
	DeleteCoffee(ctx context.Context, coffeeD domain.CoffeeID) error
}

type ChatToolUsecase struct {
	UserRepo   repository.UserRepository
	CoffeeRepo repository.CoffeeRepository
	JoinRepo   repository.JoinCoffeeToUserRepository
}

func NewChatToolUsecase(userRepo repository.UserRepository, coffeeRepo repository.CoffeeRepository, joinRepo repository.JoinCoffeeToUserRepository) *ChatToolUsecase {
	return &ChatToolUsecase{UserRepo: userRepo, CoffeeRepo: coffeeRepo, JoinRepo: joinRepo}
}

func (c *ChatToolUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	err := c.UserRepo.CreateUser(ctx, user)
	return err
}

func (c *ChatToolUsecase) GetUserByCoffeeID(ctx context.Context, coffeeID domain.CoffeeID) ([]domain.JoinCoffeeToUser, error) {
	join, err := c.JoinRepo.GetJoinByCoffeeID(ctx, coffeeID)
	return join, err
}

func (c *ChatToolUsecase) GetUserByUserID(ctx context.Context, userID domain.UserID) (domain.User, error) {
	user, err := c.UserRepo.GetUser(ctx, userID)
	return user, err
}

func (c *ChatToolUsecase) CreateCoffee(ctx context.Context, coffee *domain.Coffee) error {
	err := c.CoffeeRepo.CreateCoffee(ctx, coffee)
	return err
}

func (c *ChatToolUsecase) EditCoffeeConfig(ctx context.Context, beforeCoffee *domain.Coffee, afterCoffee *domain.Coffee) error {
	err := c.CoffeeRepo.UpdateCoffee(ctx, "", afterCoffee)
	return err
}

func (c *ChatToolUsecase) GetCoffeeByCoffeeID(ctx context.Context, coffeeID domain.CoffeeID) (domain.Coffee, error) {
	coffee, err := c.CoffeeRepo.GetCoffeeByCoffeeID(ctx, coffeeID)
	return coffee, err
}

func (c *ChatToolUsecase) DeleteCoffeeByUserIDAndCoffeeID(ctx context.Context, userID domain.UserID, coffeeID domain.CoffeeID) error {
	err := c.JoinRepo.DeleteConnectionUserIDToCoffeeID(ctx, userID, coffeeID)
	return err
}

func (c *ChatToolUsecase) CreateCoffeeConnection(ctx context.Context, joinCoffeeToUser *domain.JoinCoffeeToUser) error {
	err := c.JoinRepo.CreateConnectionUserIDToCoffeeID(ctx, joinCoffeeToUser)
	return err
}

func (c *ChatToolUsecase) DeleteCoffee(ctx context.Context, coffeeD domain.CoffeeID) error {
	err := c.CoffeeRepo.DeleteCoffee(ctx, coffeeD)
	return err
}

func (c *ChatToolUsecase) GetCoffees(ctx context.Context) ([]domain.Coffee, error) {
	coffee, err := c.CoffeeRepo.GetCoffees(ctx)
	if err != nil {
		return nil, err
	}
	return coffee, nil
}

func (c *ChatToolUsecase) GetCoffeeByUserID(ctx context.Context, userID domain.UserID) ([]domain.JoinCoffeeToUser, error) {
	joinCoffeeToUsers, err := c.JoinRepo.GetJoinByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return joinCoffeeToUsers, nil
}
