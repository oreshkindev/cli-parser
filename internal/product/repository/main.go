package repository

import (
	"cli-parser/common"
	"cli-parser/database"
	"cli-parser/internal/product/entity"
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

func (repository Repository) Find(id int) (*entity.Product, error) {

	var (
		entry entity.Product
	)

	if err = repository.connection.QueryRow(repository.context, "SELECT * FROM products WHERE id = $1", id).Scan(&entry.ID, &entry.BrandsID, &entry.Characteristics, &entry.CreatedAt, &entry.Depth, &entry.Description, &entry.DescriptionTranslated, &entry.Height, &entry.Href, &entry.MarketplacesID, &entry.Markup, &entry.Name, &entry.NameTranslated, &entry.Price, &entry.Quantity, &entry.UpdatedAt, &entry.Weight, &entry.Width); err != nil {
		return nil, common.Error(err)
	}

	return &entry, nil
}

func (repository Repository) Save(entries []entity.Product) error {

	var (
		batch pgx.Batch
	)

	for _, entry := range entries {

		batch.Queue("INSERT INTO products (id, brands_id, characteristics, depth, description, description_translated, height, href, marketplaces_id, markup, name, name_translated, price, quantity, weight, width) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name, href = EXCLUDED.href, description = EXCLUDED.description, description_translated = EXCLUDED.description_translated, price = EXCLUDED.price, quantity = EXCLUDED.quantity, weight = EXCLUDED.weight, width = EXCLUDED.width", entry.ID, entry.BrandsID, entry.Characteristics, entry.Depth, entry.Description, entry.DescriptionTranslated, entry.Height, entry.Href, entry.MarketplacesID, entry.Markup, entry.Name, entry.NameTranslated, entry.Price, entry.Quantity, entry.Weight, entry.Width)
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
