package repository

import (
	"cli-parser/common"
	"cli-parser/database"
	"cli-parser/internal/brand/entity"
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

func (repository Repository) Find(id int) (*entity.Brand, error) {

	var (
		entry entity.Brand
	)

	if err = repository.connection.QueryRow(repository.context, "SELECT id, name FROM brands WHERE id = $1", id).Scan(&entry.ID, &entry.Name); err != nil {
		return nil, common.Error(err)
	}

	return &entry, nil
}

func (repository Repository) Save(entries []entity.Brand) error {

	var (
		batch pgx.Batch
	)

	for _, entry := range entries {
		batch.Queue("INSERT INTO brands (id, name) VALUES ($1, $2) ON CONFLICT DO NOTHING", entry.ID, entry.Name)
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
