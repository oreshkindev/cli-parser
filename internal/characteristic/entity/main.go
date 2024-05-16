package entity

type (
	ResponseData[T any] struct {
		Count    int `json:"count"`
		Next     any `json:"next,omitempty"`
		Previous any `json:"previous,omitempty"`
		Results  []T `json:"results"`
	}

	Characteristic struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		NameTranslated string `json:"name_translated"`
	}

	Usecase interface {
		Find(id int) (*Characteristic, error)
		Fetch(l int, o int) ([]Characteristic, error)

		Sync() error
	}

	Repository interface {
		Find(id int) (*Characteristic, error)
		Save(entries []Characteristic) error
	}
)
