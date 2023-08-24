package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	domain "github.com/Keyl0ve/coffee-taste-app/src/domain/model"
	"github.com/Keyl0ve/coffee-taste-app/src/usecase"
)

type ServiceDriver struct {
	Controller usecase.Usecase
}

func NewServiceDriver(controller *usecase.ChatToolUsecase) *ServiceDriver {
	return &ServiceDriver{
		Controller: controller,
	}
}

// curl http://localhost:8080/api/coffee/get
func (s *ServiceDriver) CoffeeGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	coffees, err := s.Controller.GetCoffees(ctx)
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Coffees []domain.Coffee `json:"coffees"`
	}{
		Coffees: coffees,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl -H 'coffeeName:general' http://localhost:8080/api/coffee/create
func (s *ServiceDriver) CoffeeCreate(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	coffeeName := r.Header.Get("coffeeName")
	coffee := domain.NewCoffee(coffeeName, now)

	err := s.Controller.CreateCoffee(ctx, coffee)
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Success bool `json:"Success"`
	}{
		Success: true,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl -H 'userID:UserID_testUser1_12345' http://localhost:8080/api/user/get
func (s *ServiceDriver) UserGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	user, err := s.Controller.GetUserByUserID(ctx, domain.UserID(userID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		User domain.User `json:"User"`
	}{
		User: user,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl -H 'coffeeID:CoffeeID_testCoffee1_12345' -H 'userID:UserID_testUser1_12345' http://localhost:8080/api/join/delete
func (s *ServiceDriver) JoinDelete(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	coffeeID := r.Header.Get("coffeeID")

	// TODO: 既に脱退している場合、エラーを返す
	err := s.Controller.DeleteCoffeeByUserIDAndCoffeeID(ctx, domain.UserID(userID), domain.CoffeeID(coffeeID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Success bool `json:"Success"`
	}{
		Success: true,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl -H 'coffeeID:CoffeeID_testCoffee1_12345' -H 'userID:UserID_testUser1_12345' http://localhost:8080/api/join/create
func (s *ServiceDriver) JoinCreate(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	coffeeID := r.Header.Get("coffeeID")

	// userID から User 情報を取得
	user, err := s.Controller.GetUserByUserID(ctx, domain.UserID(userID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// coffeeID から Coffee 情報を取得
	coffee, err := s.Controller.GetCoffeeByCoffeeID(ctx, domain.CoffeeID(coffeeID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	if userID == "" || user.UserName == "" || coffeeID == "" || coffee.CoffeeName == "" {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("ユーザー、もしくはチャンネルが不正です。"))
	}

	join := domain.NewJoinCoffeeToUser(domain.UserID(userID), user.UserName, domain.CoffeeID(coffeeID), coffee.CoffeeName)

	// TODO: 既に参加している場合、エラーを返す
	err = s.Controller.CreateCoffeeConnection(ctx, join)
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Success bool `json:"Success"`
	}{
		Success: true,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl -H 'coffeeID:CoffeeID_testCoffee1_12345' http://localhost:8080/api/join/get/user
func (s *ServiceDriver) JoinGetUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	coffeeID := r.Header.Get("coffeeID")
	if coffeeID == "" {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", "header の情報が不足しています。"))
	}

	joinCoffeeToUsers, err := s.Controller.GetUserByCoffeeID(ctx, domain.CoffeeID(coffeeID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Join []domain.JoinCoffeeToUser `json:"join"`
	}{
		Join: joinCoffeeToUsers,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl -H 'userID:UserID_testUser1_12345' http://localhost:8080/api/join/get/coffee
func (s *ServiceDriver) JoinGetCoffee(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// UserID を header から取得
	userID := r.Header.Get("userID")

	// userID が空文字だったらエラーを返す
	if userID == "" {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", "header の情報が不足しています。"))
	}

	joinCoffeeToUsers, err := s.Controller.GetCoffeeByUserID(ctx, domain.UserID(userID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Join []domain.JoinCoffeeToUser `json:"join"`
	}{
		Join: joinCoffeeToUsers,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}
