package port

import (
	"context"

	"github.com/Keyl0ve/coffee-taste-app/backend/entity"
)

type UserInputPort interface {
	GetUserByID(ctx context.Context, userID string)
}

type UserOutputPort interface {
	Render(*entity.User)
	RenderError(error)
}

// userのCRUDに対するDB用のポート
type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
}
