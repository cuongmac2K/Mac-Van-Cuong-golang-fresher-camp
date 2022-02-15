package citiesbiz

import (
	"context"
	"demo/modules/cities/citiesmodel"
)

type CreateCitiesStore interface {
	Create(ctx context.Context, data *citiesmodel.CitiesCreate) error
}
type createCitiesBiz struct {
	store CreateCitiesStore
}

func NewCreateCitiesBiz(store CreateCitiesStore) *createCitiesBiz {
	return &createCitiesBiz{store: store}
}
func (biz *createCitiesBiz) CreateCities(ctx context.Context, data *citiesmodel.CitiesCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	err := biz.store.Create(ctx, data)
	return err
}
