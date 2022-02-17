package citiesbiz

import (
	"context"
	"demo/common"
	"demo/modules/cities/citiesmodel"
)

type UpdateCitiesStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*citiesmodel.Cities, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *citiesmodel.CitiesUpdate,
	) error
}

type updateCitiesBiz struct {
	store UpdateCitiesStore
}

func NewUpdateCitiesBiz(store UpdateCitiesStore) *updateCitiesBiz {
	return &updateCitiesBiz{store: store}
}

func (biz *updateCitiesBiz) UpdateCities(ctx context.Context, id int, data *citiesmodel.CitiesUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(citiesmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(citiesmodel.EntityName, nil)
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(citiesmodel.EntityName, err)
	}

	return nil
}
