package entity

import (
	"time"
)

type (
	ResponseData[T any] struct {
		Count    int `json:"count"`
		Next     any `json:"next,omitempty"`
		Previous any `json:"previous,omitempty"`
		Results  []T `json:"results"`
	}

	Product struct {
		ID                    int       `json:"id"`
		BrandsID              int       `json:"brand"`
		Characteristics       []int     `json:"characteristic_values"`
		CreatedAt             time.Time `json:"created_at"`
		Depth                 string    `json:"depth"`
		Description           string    `json:"description"`
		DescriptionTranslated string    `json:"description_translated"`
		Height                string    `json:"height"`
		Href                  string    `json:"url"`
		MarketplacesID        int       `json:"marketplace"`
		Markup                int       `json:"markup"`
		Name                  string    `json:"name"`
		NameTranslated        string    `json:"name_translated"`
		Price                 string    `json:"price"`
		Quantity              int       `json:"quantity"`
		UpdatedAt             time.Time `json:"updated_at"`
		Weight                string    `json:"weight"`
		Width                 string    `json:"width"`
	}

	Usecase interface {
		Find(id int) (*Product, error)
		Fetch(l int, o int) ([]Product, error)

		Sync() error
	}

	Repository interface {
		Find(id int) (*Product, error)
		Save(entries []Product) error
	}
)
