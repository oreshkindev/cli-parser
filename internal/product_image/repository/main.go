package repository

import (
	"cli-parser/common"
	"cli-parser/database"
	"cli-parser/internal/product_image/entity"
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

func (repository Repository) Find(id int) (*entity.ProductImage, error) {

	var (
		entry entity.ProductImage
	)

	if err = repository.connection.QueryRow(repository.context, "SELECT * FROM products_images WHERE id = $1", id).Scan(&entry.ID, &entry.ProductsID, &entry.Href); err != nil {
		return nil, common.Error(err)
	}

	return &entry, nil
}

func (repository Repository) Save(entries []entity.ProductImage) error {

	var (
		batch pgx.Batch
	)

	for _, entry := range entries {

		batch.Queue("INSERT INTO products_images (products_id, href) VALUES ($1, $2) ON CONFLICT DO NOTHING", entry.ProductsID, entry.Href)
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
