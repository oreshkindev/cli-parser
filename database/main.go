package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool *pgxpool.Pool
)

type (
	Database struct {
		*pgxpool.Pool
	}
)

func New(context context.Context, url string) (*Database, error) {

	var (
		err error
	)

	if pool, err = pgxpool.New(context, url); err != nil {
		return nil, err
	}

	return &Database{pool}, nil
}
