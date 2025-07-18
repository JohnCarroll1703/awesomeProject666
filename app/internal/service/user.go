package service

import (
	"awesomeProject666/app/internal/config"
	"awesomeProject666/app/internal/model"
	"awesomeProject666/app/internal/repository/postgres"
	"awesomeProject666/app/internal/schema"
	"context"
	"errors"
	"github.com/google/uuid"
)

type UserService struct {
	repo *postgres.UserRepository
	cfg  *config.Config
}

func NewUserService(repo *postgres.UserRepository,
	cfg *config.Config) *UserService {
	return &UserService{
		repo: repo,
		cfg:  cfg,
	}
}

func (u *UserService) GetAllUsers(ctx context.Context) (
	[]schema.UserResponse, error) {
	users, err := u.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	var response []schema.UserResponse
	for _, user := range users {
		response = append(response, schema.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return response, nil
}

func (u *UserService) CreateUser(ctx context.Context, req schema.CreateUserRequest) error {
	user := model.User{
		ID:       uuid.New(),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	return u.repo.CreateUser(ctx, &user)
}

func (u *UserService) GetUser(ctx context.Context, filter schema.UserFilters) (
	[]schema.UserResponse, error) {
	users, err := u.repo.GetUser(ctx, filter)
	if err != nil {
		return nil, err
	}

	var response []schema.UserResponse
	for _, user := range users {
		response = append(response, schema.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}

	return response, nil
}

func (u *UserService) GetByEmailAndPassword(ctx context.Context, email, password string) (*model.User, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return &model.User{}, err
	}

	// Безопасная проверка — можно bcrypt позже прикрутить
	if user.Password != password {
		return &model.User{}, errors.New("invalid password")
	}

	return user, nil
}
