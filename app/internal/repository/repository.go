package repository

import (
	"awesomeProject666/app/internal/model"
	"awesomeProject666/app/internal/repository/postgres"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	UserRepo *postgres.UserRepository
	ItemRepo *postgres.ItemRepository
}

type User interface {
	GetAllUsers(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, u *model.User) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type Item interface {
	GetAllItems(ctx context.Context) ([]model.Item, error)
}

func NewRepository(db *pgxpool.Pool) *Repositories {
	userRepo := postgres.NewUser(db)
	itemRepo := postgres.NewItem(db)
	return &Repositories{
		UserRepo: userRepo,
		ItemRepo: itemRepo,
	}
}
