package characteristic_extended

import (
	"cli-parser/database"
	"cli-parser/internal/characteristic_extended/entity"
	"cli-parser/internal/characteristic_extended/repository"
	"cli-parser/internal/characteristic_extended/usecase"
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
