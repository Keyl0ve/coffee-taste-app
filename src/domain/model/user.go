package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserID string

func NewUserID(name string) UserID {
	uuid, _ := uuid.NewRandom()
	return UserID(fmt.Sprintf("userID_%s_%s", name, uuid))
}

type User struct {
	UserID    UserID
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
