package model

type JoinCoffeeToUser struct {
	UserID     UserID
	UserName   string
	CoffeeID   CoffeeID
	CoffeeName string
}

func NewJoinCoffeeToUser(userID UserID, userName string, coffeeID CoffeeID, coffeeName string) *JoinCoffeeToUser {
	return &JoinCoffeeToUser{
		UserID:     userID,
		UserName:   userName,
		CoffeeID:   coffeeID,
		CoffeeName: coffeeName,
	}
}
