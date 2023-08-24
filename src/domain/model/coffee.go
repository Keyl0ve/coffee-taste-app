package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CoffeeID string

func NewCoffeeID(name string) CoffeeID {
	uuid, _ := uuid.NewRandom()
	return CoffeeID(fmt.Sprintf("CoffeeID_%s_%s", name, uuid))
}

type Coffee struct {
	CoffeeID   CoffeeID
	CoffeeName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewCoffee(name string, now time.Time) *Coffee {
	return &Coffee{
		CoffeeID:   NewCoffeeID(name),
		CoffeeName: name,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
