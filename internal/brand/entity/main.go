package entity

type (
	ResponseData[T any] struct {
		Count    int `json:"count"`
		Next     any `json:"next,omitempty"`
		Previous any `json:"previous,omitempty"`
		Results  []T `json:"results"`
	}

	Brand struct {
		ID   int
		Name string `json:"name"`
	}

	Usecase interface {
		Find(id int) (*Brand, error)
		Fetch(l int, o int) ([]Brand, error)

		Sync() error
	}

	Repository interface {
		Find(id int) (*Brand, error)
		Save(entries []Brand) error
	}
)
