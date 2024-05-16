package entity

type (
	ResponseData[T any] struct {
		Count    int `json:"count"`
		Next     any `json:"next,omitempty"`
		Previous any `json:"previous,omitempty"`
		Results  []T `json:"results"`
	}

	CharacteristicExtended struct {
		ID               int    `json:"id"`
		CharacteristicID int    `json:"characteristic"`
		Name             string `json:"value"`
		NameTranslated   string `json:"name_translated"`
	}

	Usecase interface {
		Find(id int) (*CharacteristicExtended, error)
		Fetch(l int, o int) ([]CharacteristicExtended, error)

		Sync() error
	}

	Repository interface {
		Find(id int) (*CharacteristicExtended, error)
		Save(entries []CharacteristicExtended) error
	}
)
