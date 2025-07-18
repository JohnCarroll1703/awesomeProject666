package service

import (
	"awesomeProject666/app/internal/config"
	"awesomeProject666/app/internal/repository"
	"awesomeProject666/app/internal/schema"
	"context"
)

type User interface {
	GetAllUsers(ctx context.Context) ([]schema.UserResponse, error)
	CreateUser(ctx context.Context, req schema.CreateUserRequest) error
}

type Item interface {
	GetAllItems(ctx context.Context) ([]schema.ItemResponse, error)
}

type Services struct {
	UserService *UserService
	ItemService *ItemService
}

func NewServices(repo *repository.Repositories,
	cfg *config.Config) *Services {
	return &Services{
		UserService: NewUserService(repo.UserRepo, cfg),
		ItemService: NewItemService(repo.ItemRepo, cfg),
	}
}
