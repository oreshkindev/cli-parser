package internal

import (
	"cli-parser/database"
	"cli-parser/internal/brand"
	"cli-parser/internal/characteristic"
	"cli-parser/internal/characteristic_extended"
	"context"
	"os"

	"github.com/go-resty/resty/v2"
)

type (
	Manager struct {
		Brand                  brand.Manager
		Characteristic         characteristic.Manager
		CharacteristicExtended characteristic_extended.Manager
	}
)

func New(context context.Context, connection *database.Database) *Manager {

	http := resty.New()

	http.SetHeader("Content-Type", "application/json")

	http.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	http.SetBaseURL("https://parser.standard-it.ru/api/v1/provider/")

	return &Manager{
		Brand:                  *brand.New(context, connection, http),
		Characteristic:         *characteristic.New(context, connection, http),
		CharacteristicExtended: *characteristic_extended.New(context, connection, http),
	}
}
