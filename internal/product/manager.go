package product

import (
	"cli-parser/database"
	"cli-parser/internal/product/entity"
	"cli-parser/internal/product/repository"
	"cli-parser/internal/product/usecase"
	"context"

	"github.com/go-resty/resty/v2"
)

type (
	Manager struct {
		entity.Usecase
	}
)

func New(context context.Context, connection *database.Database, http *resty.Client) *Manager {

	repository := repository.New(context, connection)

	return &Manager{
		usecase.New(context, http, repository),
	}
}
