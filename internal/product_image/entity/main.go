package entity

type (
	ResponseData[T any] struct {
		Count    int `json:"count"`
		Next     any `json:"next,omitempty"`
		Previous any `json:"previous,omitempty"`
		Results  []T `json:"results"`
	}

	ProductImage struct {
		ID         int
		ProductsID int      `json:"product"`
		Href       []string `json:"href"`
		Image      string   `json:"image"`
	}

	Usecase interface {
		Find(id int) (*ProductImage, error)
		Fetch(l int, o int) ([]ProductImage, error)

		Sync() error
	}

	Repository interface {
		Find(id int) (*ProductImage, error)
		Save(entries []ProductImage) error
	}
)
