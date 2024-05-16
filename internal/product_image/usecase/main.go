package usecase

import (
	"cli-parser/internal/product_image/entity"
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

		entries []entity.ProductImage

		groupedHref = make(map[int][]string)
	)

	for {
		if entries, err = usecase.Fetch(l, o); err != nil {
			return err
		}

		if len(entries) == 0 {
			break
		}

		for _, entry := range entries {
			groupedHref[entry.ProductsID] = append(groupedHref[entry.ProductsID], entry.Image)
		}

		o += l
	}

	for productID, href := range groupedHref {
		entries = append(entries, entity.ProductImage{
			ProductsID: productID,
			Href:       href,
		})
	}

	if err = usecase.Save(entries); err != nil {
		return err
	}

	return nil
}

func (usecase Usecase) Fetch(l, o int) ([]entity.ProductImage, error) {

	var (
		responseData entity.ResponseData[entity.ProductImage]
	)

	request := usecase.http.R().SetQueryParams(map[string]string{
		"limit":  strconv.Itoa(l),
		"offset": strconv.Itoa(o),
	})

	if response, err = request.Get("product_images/"); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(response.Body(), &responseData); err != nil {
		return nil, err
	}

	return responseData.Results, nil
}

func (usecase Usecase) Find(id int) (*entity.ProductImage, error) {

	return usecase.repository.Find(id)
}

func (usecase Usecase) Save(entries []entity.ProductImage) error {

	return usecase.repository.Save(entries)
}
