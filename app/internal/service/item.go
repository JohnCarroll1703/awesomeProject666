package service

import (
	"awesomeProject666/app/internal/config"
	"awesomeProject666/app/internal/repository/postgres"
	"awesomeProject666/app/internal/schema"
	"context"
)

type ItemService struct {
	repo *postgres.ItemRepository
	cfg  *config.Config
}

func NewItemService(repo *postgres.ItemRepository, cfg *config.Config) *ItemService {
	return &ItemService{repo: repo, cfg: cfg}
}

func (i *ItemService) GetAllItems(ctx context.Context) (
	[]schema.ItemResponse, error) {
	res, err := i.repo.GetItems(ctx)
	if err != nil {
		return nil, err
	}

	var response []schema.ItemResponse
	for _, item := range res {
		response = append(response, schema.ItemResponse{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Stock:       item.Stock,
			CategoryID:  item.CategoryID,
			CountryID:   item.CountryID,
			AddedAt:     item.AddedAt,
		})
	}
	return response, nil
}
