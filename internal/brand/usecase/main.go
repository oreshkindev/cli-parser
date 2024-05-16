package usecase

import (
	"cli-parser/internal/brand/entity"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-resty/resty/v2"
)

type (
	Usecase struct {
		context    context.Context
		http       *resty.Client
		repository entity.Repository
	}
)

var (
	err      error
	response *resty.Response
)

func New(context context.Context, http *resty.Client, repository entity.Repository) *Usecase {
	return &Usecase{
		context:    context,
		http:       http,
		repository: repository,
	}
}

func (usecase *Usecase) Sync() error {

	var (
		l = 250
		o = 0

		entries []entity.Brand
	)

	for {
		if entries, err = usecase.Fetch(l, o); err != nil || len(entries) == 0 {
			break
		}

		if err = usecase.Save(entries); err != nil {
			return err
		}

		o += l

	}

	return nil
}

func (usecase Usecase) Fetch(l, o int) ([]entity.Brand, error) {

	var (
		responseData entity.ResponseData[entity.Brand]
	)

	request := usecase.http.R().SetQueryParams(map[string]string{
		"limit":  strconv.Itoa(l),
		"offset": strconv.Itoa(o),
	})

	if response, err = request.Get("brands/"); err != nil {

		return nil, err
	}

	if err = json.Unmarshal(response.Body(), &responseData); err != nil {
		return nil, err
	}

	return responseData.Results, nil
}

func (usecase Usecase) Find(id int) (*entity.Brand, error) {

	return usecase.repository.Find(id)
}

func (usecase Usecase) Save(entries []entity.Brand) error {

	return usecase.repository.Save(entries)
}
