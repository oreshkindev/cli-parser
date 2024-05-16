package brand

import (
	"cli-parser/database"
	"cli-parser/internal/brand/entity"
	"cli-parser/internal/brand/repository"
	"cli-parser/internal/brand/usecase"
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
