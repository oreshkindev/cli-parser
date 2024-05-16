package repository

import (
	"cli-parser/common"
	"cli-parser/database"
	"cli-parser/internal/characteristic/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

type (
	Repository struct {
		context    context.Context
		connection *database.Database
	}
)

var (
	err error
)

func New(context context.Context, connection *database.Database) *Repository {
	return &Repository{
		context:    context,
		connection: connection,
	}
}

func (repository Repository) Find(id int) (*entity.Characteristic, error) {

	var (
		entry entity.Characteristic
	)

	if err = repository.connection.QueryRow(repository.context, "SELECT name FROM characteristics WHERE id = $1", id).Scan(&entry.ID, &entry.Name); err != nil {
		return nil, common.Error(err)
	}

	return &entry, nil
}

func (repository Repository) Save(entries []entity.Characteristic) error {

	var (
		batch pgx.Batch
	)

	for _, entry := range entries {
		batch.Queue("INSERT INTO characteristics (id, name) VALUES ($1, $2) ON CONFLICT DO NOTHING", entry.ID, entry.Name)
	}

	results := repository.connection.SendBatch(repository.context, &batch)

	defer results.Close()

	for range entries {
		if _, err := results.Exec(); err != nil {
			return common.Error(err)
		}
	}

	return nil
}
