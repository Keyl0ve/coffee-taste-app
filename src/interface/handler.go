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

// curl -H 'userID:1' http://localhost:8080/api/user/get
func (s *ServiceDriver) UserGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	users, err := s.Controller.GetUsers(ctx)
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		Users []domain.User `json:"users"`
	}{
		Users: users,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

// curl -H 'userName:kyo' -H 'password:password' http://localhost:8080/api/user/create
func (s *ServiceDriver) UserCreate(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	userName := r.Header.Get("userName")
	password := r.Header.Get("password")
	user := domain.NewUser(userName, password, now)

	err := s.Controller.CreateUser(ctx, user)
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
