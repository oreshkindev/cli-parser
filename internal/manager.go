package internal

import (
	"cli-parser/database"
	"cli-parser/internal/brand"
	"context"
	"os"

	"github.com/go-resty/resty/v2"
)

type (
	Manager struct {
		Brand brand.Manager
	}
)

func New(context context.Context, connection *database.Database) *Manager {

	http := resty.New()

	http.SetHeader("Content-Type", "application/json")

	http.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	http.SetBaseURL("https://parser.standard-it.ru/api/v1/provider/")

	return &Manager{
		Brand: *brand.New(context, connection, http),
	}
}
