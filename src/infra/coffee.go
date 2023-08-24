package infra

import (
	"context"
	"database/sql"
	"log"

	domain "github.com/Keyl0ve/coffee-taste-app/src/domain/model"
)

type CoffeeRepository struct {
	Conn *sql.DB
}

func NewCoffeeRepository(conn *sql.DB) *CoffeeRepository {
	return &CoffeeRepository{Conn: conn}
}

func ScanCoffees(rows *sql.Rows) ([]domain.Coffee, int, error) {
	coffees := make([]domain.Coffee, 0)

	for rows.Next() {
		var v domain.Coffee
		if err := rows.Scan(&v.CoffeeID, &v.CoffeeName, &v.CreatedAt, &v.UpdatedAt); err != nil {
			log.Printf("[ERROR] scan ScanCoffees: %+v", err)
			return nil, 0, err
		}
		coffees = append(coffees, v)
	}

	return coffees, len(coffees), nil
}

func (c CoffeeRepository) CreateCoffee(ctx context.Context, coffee *domain.Coffee) error {
	query := "INSERT INTO coffee (coffee_ID, coffee_name, created_at, updated_at) VALUES (?,?,?,?) "
	_, err := c.Conn.ExecContext(ctx, query, coffee.CoffeeID, coffee.CoffeeName, coffee.CreatedAt, coffee.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR] can't create CreateCoffee: %+v", err)
		return nil
	}

	return nil
}

func (c CoffeeRepository) GetCoffees(ctx context.Context) ([]domain.Coffee, error) {
	query := "SELECT * FROM coffee"
	rows, err := c.Conn.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[ERROR] not found Coffees: %+v", err)
		return nil, err
	}

	coffees, _, err := ScanCoffees(rows)
	if err != nil {
		log.Printf("[ERROR] can't scan Coffees: %+v", err)
		return nil, err
	}

	return coffees, nil
}

func (c CoffeeRepository) GetCoffeeByCoffeeID(ctx context.Context, coffeeID domain.CoffeeID) (domain.Coffee, error) {
	query := "SELECT * FROM	coffee WHERE coffee_id = ?"
	rows, err := c.Conn.QueryContext(ctx, query, coffeeID)
	if err != nil {
		log.Printf("[ERROR] not found Coffees: %+v", err)
		return domain.Coffee{}, err
	}

	var coffee domain.Coffee
	for rows.Next() {
		if err := rows.Scan(&coffee.CoffeeID, &coffee.CoffeeName, &coffee.CreatedAt, &coffee.UpdatedAt); err != nil {
			log.Printf("[ERROR] scan ScanCoffees: %+v", err)
			return domain.Coffee{}, err
		}
	}
	return coffee, nil
}

func (c CoffeeRepository) UpdateCoffee(ctx context.Context, coffeeID domain.CoffeeID, updatedCoffee *domain.Coffee) error {
	query := "UPDATE coffee set CoffeeName = ? WHERE CoffeeID = ? "
	_, err := c.Conn.ExecContext(ctx, query, updatedCoffee, coffeeID)
	if err != nil {
		log.Printf("[ERROR] can't UpdateCoffee: %+v", err)
		return nil
	}

	return nil
}

func (c CoffeeRepository) DeleteCoffee(ctx context.Context, coffeeID domain.CoffeeID) error {
	query := "DELETE FROM coffee WHERE id = ?"
	_, err := c.Conn.ExecContext(ctx, query, coffeeID)
	if err != nil {
		log.Printf("[ERROR] can't DeleteCoffee: %+v", err)
		return nil
	}

	return nil
}
